package taskhandler

import (
	"context"
	"strings"
	"time"

	apiv1 "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func (th *TaskHandler) syncTransactions(ctx context.Context, userId uint64, link *apiv1.Link, accessToken, trigger string) error {
	var (
		postgresClient   = th.postgresDb
		clickhouseClient = th.clickhouseDb
		plaidClient      = th.plaidClient
		cursor           *string
	)

	// get the last plaid sync
	lasPlaidSync := link.GetPlaidSync()
	if lasPlaidSync != nil {
		cursor = &lasPlaidSync.NextCursor
	}

	syncResult, err := plaidClient.Sync(ctx, cursor, &accessToken)
	if err != nil {
		return errors.Wrap(err, "failed to sync with plaid")
	}

	// record the sync event in the database
	linkId := link.Id
	nextCursor := syncResult.NextCursor
	added := len(syncResult.New)
	deleted := len(syncResult.Deleted)
	modified := len(syncResult.Updated)
	if err := postgresClient.RecordPlaidSync(ctx, userId, linkId, trigger, nextCursor, int64(added), int64(modified), int64(deleted)); err != nil {
		return err
	}

	// If we received nothing to insert/update/remove then do nothing
	if len(syncResult.New)+len(syncResult.Updated)+len(syncResult.Deleted) == 0 {
		log.Info("no new data from plaid, nothing to be done")
		return nil
	}

	plaidTransactions := append(syncResult.New, syncResult.Updated...)
	th.logger.Info("received new transactions from plaid", zap.Int("count", len(syncResult.New)))

	plaidTransactionIds := make([]string, 0, len(syncResult.Updated))
	for _, transaction := range syncResult.Updated {
		plaidTransactionIds = append(plaidTransactionIds, transaction.GetTransactionId())
	}

	// query the database for all of the transactions that we already have
	plaidTxnIdToTxnMap := make(map[string]*apiv1.Transaction, 0)
	txnFound, err := clickhouseClient.GetTransactionsByPlaidTransactionIds(ctx, plaidTransactionIds)
	if err != nil {
		return err
	}

	for _, txn := range txnFound {
		plaidTxnIdToTxnMap[txn.GetTransactionId()] = txn
	}

	transactionsToUpdate := make([]*apiv1.Transaction, 0, len(plaidTransactions))
	transactionsToInsert := make([]*apiv1.Transaction, 0, len(plaidTransactions))
	for _, plaidTransaction := range plaidTransactions {
		amount := plaidTransaction.GetAmount()
		transactionName := plaidTransaction.GetName()
		// We only want to make the transaction name be the merchant name if the merchant name is shorter. This is
		// due to something I observed with a dominos transaction, where the merchant was improperly parsed and the
		// transaction ended up being called `Mnuslindstrom` rather than `Domino's`. This should fix that problem.
		if plaidTransaction.GetMerchantName() != "" && len(plaidTransaction.GetMerchantName()) < len(transactionName) {
			transactionName = plaidTransaction.GetMerchantName()
		}

		// if we have no present transactions in the database for the user of interest
		// then we know that we need to insert all of the transactions
		if len(txnFound) == 0 {
			plaidTransaction.UserId = userId
			plaidTransaction.LinkId = linkId
			plaidTransaction.Sign = 1
			plaidTransaction.Name = transactionName

			transactionsToInsert = append(transactionsToInsert, plaidTransaction)
			continue
		}

		existingTransaction, ok := plaidTxnIdToTxnMap[plaidTransaction.GetTransactionId()]
		if !ok {
			plaidTransaction.UserId = userId
			plaidTransaction.LinkId = linkId
			plaidTransaction.Sign = 1
			plaidTransaction.Name = transactionName

			transactionsToInsert = append(transactionsToInsert, plaidTransaction)
			continue
		} else {
			// due to clickhouse's collapsing merge tree deduplication logic, we need to
			// also append the old transactions that was updated but with a sign bit that is negative
			/**
			CollapsingMergeTree: This engine collapses rows with the same sorting key during the merge operation based on the sign column.

			To insert a row: sign = 1
			To delete a row: sign = -1
			*/
			var shouldUpdate bool = false

			var newTransaction = existingTransaction
			newTransaction.Sign = 1
			if newTransaction.Amount != amount {
				newTransaction.Amount = amount
				shouldUpdate = true
			}

			if newTransaction.Pending != plaidTransaction.GetPending() {
				newTransaction.Pending = plaidTransaction.GetPending()
				shouldUpdate = true
			}

			if strings.EqualFold(newTransaction.PendingTransactionId, plaidTransaction.GetPendingTransactionId()) {
				newTransaction.PendingTransactionId = plaidTransaction.GetPendingTransactionId()
				shouldUpdate = true
			}

			if shouldUpdate {
				transactionsToUpdate = append(transactionsToUpdate, newTransaction)
			}
		}

	}

	if len(transactionsToUpdate) > 0 {
		th.logger.Info("updating transactions", zap.Int("count", len(transactionsToUpdate)))
		if err = clickhouseClient.UpdateTransactions(ctx, &userId, transactionsToUpdate); err != nil {
			return err
		}
	}

	if len(transactionsToInsert) > 0 {
		th.logger.Info("updating transactions", zap.Int("count", len(transactionsToInsert)))
		if err := clickhouseClient.AddTransactions(ctx, &userId, transactionsToInsert); err != nil {
			th.logger.Error("failed to insert transactions", zap.Error(err))
			return err
		}
	}

	if len(syncResult.Deleted) > 0 { // Handle removed transactions
		th.logger.Info("handling removed transactions", zap.Int("count", len(syncResult.Deleted)))

		// TODO: also update the transactions in s3
		transactions, err := clickhouseClient.GetTransactionsByPlaidTransactionIds(ctx, syncResult.Deleted)
		if err != nil {
			th.logger.Error("failed to retrieve transactions by plaid transaction Id for removal", zap.Error(err))
			return err
		}

		if len(transactions) == 0 {
			log.Warnf("no transactions retrieved, nothing to be done. transactions might already have been deleted")
			return nil
		}

		if len(transactions) != len(syncResult.Deleted) {
			th.logger.Warn("number of transactions retrieved does not match expected number of transactions", zap.Int("expected", len(syncResult.Deleted)), zap.Int("found", len(transactions)))
		}

		// TODO: process goals and forecasts here
		for _, transaction := range transactions {
			if err = clickhouseClient.DeleteTransaction(ctx, &transaction.Id); err != nil {
				th.logger.Error("failed to delete transaction", zap.String("transactionId", transaction.TransactionId), zap.Error(err))
				return err
			}
		}

		th.logger.Info("successfully removed transactions", zap.Int("count", len(transactions)))
	}

	linkWasSetup := false

	// If the link status is not setup or pending expiration. Then change the status to setup
	switch link.LinkStatus {
	case apiv1.LinkStatus_LINK_STATUS_SETUP, apiv1.LinkStatus_LINK_STATUS_PENDING_EXPIRATION:
	default:
		link.LinkStatus = apiv1.LinkStatus_LINK_STATUS_SETUP
		linkWasSetup = true
	}
	link.LastSuccessfulUpdate = time.Now().String()
	if err = postgresClient.UpdateLink(ctx, link); err != nil {
		th.logger.Error("failed to update link after transaction sync", zap.Error(err))
		return err
	}

	if linkWasSetup {
		// TODO: Send the notification that the link has been set up.
		th.logger.Info("sending link setup notification")
	}

	return nil
}
