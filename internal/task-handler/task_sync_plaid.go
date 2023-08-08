package taskhandler

import (
	"context"
	"fmt"
	"strings"
	"time"

	"encoding/json"

	"github.com/SimifiniiCTO/asynq"
	apiv1 "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SyncPlaidTaskPayload struct {
	UserId      uint64 `json:"user_id"`
	AccessToken string `json:"access_token"`
	LinkId      uint64 `json:"plaid_link_item_id"`
}

func (t *SyncPlaidTaskPayload) String() *string {
	str := fmt.Sprintf("user_id: %d, access_token: %s, plaid_link_item_id: %d", t.UserId, t.AccessToken, t.LinkId)
	return &str
}

// This function creates a new asynchronous task for syncing Plaid  with the provided user
// ID and access token.
func NewSyncPlaidTask(userId uint64, accessToken string, linkId uint64) (*asynq.Task, error) {
	payload, err := json.Marshal(&SyncPlaidTaskPayload{
		UserId:      userId,
		AccessToken: accessToken,
		LinkId:      linkId,
	})

	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TaskSyncTransactions.String(), payload), nil
}

// This is a method of the `TaskHandler` struct that handles the execution of a task for syncing Plaid
// transactions. It takes in a context and a task object as parameters and returns an error. Inside the
// method, it unmarshals the payload of the task, queries Plaid for the user's transactions, sanitizes
// and transforms the transactions, inserts them into a database, and sends the data to a data
// warehouse.
func (th *TaskHandler) RunSyncPlaidTransactionsTask(ctx context.Context, task *asynq.Task) error {
	var (
		payload SyncPlaidTaskPayload
	)

	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return err
	}

	trigger := payload.String()
	accts, err := th.processSyncOperation(ctx, payload.UserId, payload.LinkId, payload.AccessToken, *trigger)
	if err != nil {
		return err
	}

	if len(accts) == 0 {
		th.logger.Warn("no accounts found for user", zap.Uint64("user_id", payload.UserId))
		return nil
	}

	return nil
}

// TODO: Refactor this piece of code
func (th *TaskHandler) processSyncOperation(ctx context.Context, userId, linkId uint64, accessToken, trigger string) ([]*apiv1.BankAccount, error) {
	var (
		postgresClient   = th.postgresDb
		clickhouseClient = th.clickhouseDb
		plaidClient      = th.plaidClient
	)

	// TODO: from the linked account we need to be able to see all accounts that the user may have chosen to deselect from access to us and remove them
	// we need explicit support for this

	// query the link from the database
	link, err := postgresClient.GetLink(ctx, userId, linkId, false)
	if err != nil {
		return nil, err
	}

	if link.PlaidLink == nil {
		th.logger.Warn("provided link does not have any plaid credentials")
		return nil, errors.New("plaid link not found")
	}

	// construct account balance objects from the synced accounts
	accountBalances := make([]*apiv1.AccountBalanceHistory, 0)

	// query plaid for all of the user's accounts
	syncBankAccounts, err := plaidClient.GetAccounts(ctx, accessToken, userId)
	if err != nil {
		return nil, err
	}

	// we only perform account sync cross referencing if the
	// number of synced accounts is greater than 0
	if len(syncBankAccounts) == 0 {
		th.logger.Warn("no accounts found for user", zap.Uint64("user_id", userId))
		return nil, nil
	} else {
		if err := th.syncBankAccountsHelper(ctx, link, syncBankAccounts); err != nil {
			return nil, err
		}

		th.logger.Info("adding acocunt balances")
		for _, account := range syncBankAccounts {
			accountBalance := &apiv1.AccountBalanceHistory{
				Time:            timestamppb.Now(),
				AccountId:       account.PlaidAccountId,
				IsoCurrencyCode: account.Currency,
				Balance:         float64(account.Balance),
				UserId:          userId,
				Sign:            1,
				Id:              "",
			}

			accountBalances = append(accountBalances, accountBalance)
		}
	}

	plaidAccountIds := make([]string, 0, len(syncBankAccounts))
	for _, plaidBankAccount := range syncBankAccounts {
		if plaidBankAccount.Subtype == "checking" || plaidBankAccount.Subtype == "savings" || plaidBankAccount.Subtype == "cd" || plaidBankAccount.Subtype == "money market" ||
			plaidBankAccount.Subtype == "gic" || plaidBankAccount.Subtype == "prepaid" || plaidBankAccount.Subtype == "health" || plaidBankAccount.Subtype == "cash management" {
			plaidAccountIds = append(plaidAccountIds, plaidBankAccount.PlaidAccountId)
		}
	}

	// // lets sync recurring transactions
	// if err := th.queryAndStoreRecurringTransactions(ctx, accessToken, userId, linkId, plaidAccountIds); err != nil {
	// 	th.logger.Error("failed to sync recurring transactions", zap.Error(err))
	// }

	// here we pull all investment accounts if authorized against the provided accesstoken
	// get the investment account ids
	newlySyncedInvestmentAccounts, err := th.syncInvestmentAccountsHelper(ctx, userId, accessToken, link)
	if err != nil {
		th.logger.Error("failed to sync investment accounts .. possibily because no investment account is tied to the following lin", zap.Error(err))
	}

	// add to the balance history
	for _, account := range newlySyncedInvestmentAccounts {
		accountBalance := &apiv1.AccountBalanceHistory{
			AccountId: account.PlaidAccountId,
			Balance:   float64(account.Balance),
			UserId:    userId,
			Sign:      1,
			Id:        "",
			Time:      timestamppb.Now(),
		}

		accountBalances = append(accountBalances, accountBalance)
	}

	// link contains liability accounts hence sync them
	// get the credit accounts
	// TODO: need to figure out wether to update the existing credit account or add a new one
	creditAccountsSet, err := plaidClient.GetPlaidLiabilityAccounts(ctx, &accessToken)
	th.logger.Info("credit accounts", zap.Any("creditAccountsSet", creditAccountsSet))
	if err != nil {
		th.logger.Error("failed to get credit accounts", zap.Error(err))
	} else {
		if len(creditAccountsSet.CrediCardAccounts) > 0 {
			syncedCreditAccounts := creditAccountsSet.CrediCardAccounts
			if err := th.syncCreditAccountsHelper(ctx, link, syncedCreditAccounts); err != nil {
				return nil, err
			}

			for _, account := range syncedCreditAccounts {
				accountBalance := &apiv1.AccountBalanceHistory{
					AccountId: account.PlaidAccountId,
					Balance:   float64(account.Balance),
					UserId:    userId,
					Sign:      1,
					Id:        "",
					Time:      timestamppb.Now(),
				}

				accountBalances = append(accountBalances, accountBalance)
			}
		}

		if len(creditAccountsSet.MortgageLoanAccts) > 0 {
			syncedMortgageLoanAccounts := creditAccountsSet.MortgageLoanAccts
			if err := th.syncMortgageAccountsHelper(ctx, link, syncedMortgageLoanAccounts); err != nil {
				return nil, err
			}

			for _, account := range syncedMortgageLoanAccounts {
				accountBalance := &apiv1.AccountBalanceHistory{
					AccountId: account.PlaidAccountId,
					Balance:   float64(account.OutstandingPrincipalBalance),
					UserId:    userId,
					Sign:      1,
					Id:        "",
					Time:      timestamppb.Now(),
				}

				accountBalances = append(accountBalances, accountBalance)
			}
		}

		if len(creditAccountsSet.StudentLoanAccts) > 0 {
			syncedStudentLoanAccounts := creditAccountsSet.StudentLoanAccts
			if err := th.syncStudentLoanAccountsHelper(ctx, link, syncedStudentLoanAccounts); err != nil {
				return nil, err
			}

			for _, account := range syncedStudentLoanAccounts {
				accountBalance := &apiv1.AccountBalanceHistory{
					AccountId: account.PlaidAccountId,
					Balance:   float64(account.GetOriginationPrincipalAmount()),
					UserId:    userId,
					Sign:      1,
					Id:        "",
					Time:      timestamppb.Now(),
				}

				accountBalances = append(accountBalances, accountBalance)
			}
		}
	}

	// update the account balances in the database
	if len(accountBalances) > 0 {
		th.logger.Info("updating account balances", zap.Int("count", len(accountBalances)))

		if err := clickhouseClient.AddAccountBalances(ctx, &userId, accountBalances); err != nil {
			th.logger.Error("failed to add account balances", zap.Error(err))
			return nil, err
		}
	}

	var cursor *string
	// get the last plaid sync

	lasPlaidSync := link.GetPlaidSync()
	if lasPlaidSync != nil {
		cursor = &lasPlaidSync.NextCursor
	}

	syncResult, err := plaidClient.Sync(ctx, cursor, &accessToken)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sync with plaid")
	}

	// record the sync event in the database
	linkId = link.Id
	nextCursor := syncResult.NextCursor
	added := len(syncResult.New)
	deleted := len(syncResult.Deleted)
	modified := len(syncResult.Updated)
	if err := postgresClient.RecordPlaidSync(ctx, userId, linkId, trigger, nextCursor, int64(added), int64(modified), int64(deleted)); err != nil {
		return nil, err
	}

	// If we received nothing to insert/update/remove then do nothing
	if len(syncResult.New)+len(syncResult.Updated)+len(syncResult.Deleted) == 0 {
		log.Info("no new data from plaid, nothing to be done")
		return nil, nil
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
		return nil, err
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
			return nil, err
		}
	}

	if len(transactionsToInsert) > 0 {
		th.logger.Info("updating transactions", zap.Int("count", len(transactionsToInsert)))
		if err := clickhouseClient.AddTransactions(ctx, &userId, transactionsToInsert); err != nil {
			th.logger.Error("failed to insert transactions", zap.Error(err))
			return nil, err
		}
	}

	if len(syncResult.Deleted) > 0 { // Handle removed transactions
		th.logger.Info("handling removed transactions", zap.Int("count", len(syncResult.Deleted)))

		// TODO: also update the transactions in s3
		transactions, err := clickhouseClient.GetTransactionsByPlaidTransactionIds(ctx, syncResult.Deleted)
		if err != nil {
			th.logger.Error("failed to retrieve transactions by plaid transaction Id for removal", zap.Error(err))
			return nil, err
		}

		if len(transactions) == 0 {
			log.Warnf("no transactions retrieved, nothing to be done. transactions might already have been deleted")
			return nil, nil
		}

		if len(transactions) != len(syncResult.Deleted) {
			th.logger.Warn("number of transactions retrieved does not match expected number of transactions", zap.Int("expected", len(syncResult.Deleted)), zap.Int("found", len(transactions)))
		}

		// TODO: process goals and forecasts here
		for _, transaction := range transactions {
			if err = clickhouseClient.DeleteTransaction(ctx, &transaction.Id); err != nil {
				th.logger.Error("failed to delete transaction", zap.String("transactionId", transaction.TransactionId), zap.Error(err))
				return nil, err
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
		return nil, err
	}

	if linkWasSetup {
		// TODO: Send the notification that the link has been set up.
		th.logger.Info("sending link setup notification")
	}

	return syncBankAccounts, nil
}
