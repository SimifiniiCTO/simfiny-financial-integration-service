package taskhandler

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/SimifiniiCTO/asynq"
	apiv1 "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type PullTransactionsTaskPayload struct {
	UserId      uint64    `json:"user_id"`
	LinkId      uint64    `json:"link_id"`
	AccessToken string    `json:"access_token"`
	Start       time.Time `json:"start"`
	End         time.Time `json:"end"`
}

func NewPullTransactionsTask(userId, linkId uint64, accessToken string, startTime, endTime time.Time) (*asynq.Task, error) {
	rec := &PullTransactionsTaskPayload{
		UserId:      userId,
		LinkId:      linkId,
		AccessToken: accessToken,
		Start:       startTime,
		End:         endTime,
	}

	payload, err := json.Marshal(rec)
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TaskPullTransactions.String(), payload), nil
}

func (th *TaskHandler) RunPullTransactionsTask(ctx context.Context, task *asynq.Task) error {
	var (
		postgresClient   = th.postgresDb
		clickhouseClient = th.clickhouseDb
		plaidClient      = th.plaidClient
		payload          PullTransactionsTaskPayload
	)

	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return err
	}

	userId := payload.UserId
	linkId := payload.LinkId
	accessToken := payload.AccessToken
	startTime := payload.Start
	endTime := payload.End

	// query the link from the database
	link, err := postgresClient.GetLink(ctx, userId, linkId, false)
	if err != nil {
		return err
	}

	if link.PlaidLink == nil {
		th.logger.Warn("provided link does not have any plaid credentials")
		return nil
	}

	if len(link.BankAccounts) == 0 {
		th.logger.Warn("provided link does not have any bank accounts")
		return nil
	}

	// get all the plaid bank accts
	plaidBankAccounts, err := plaidClient.GetAccounts(ctx, accessToken, userId)
	if err != nil {
		return err
	}

	if len(plaidBankAccounts) == 0 {
		th.logger.Warn("no plaid bank accounts found. nothing to sync")
		return nil
	}

	plaidIdsToBankIds := map[string]uint64{}
	plaidBankToLocalBank := map[string]*apiv1.BankAccount{}
	bankAccountIds := make([]string, 0, len(link.BankAccounts))
	for _, bankAccount := range link.BankAccounts {
		for _, plaidBankAccount := range plaidBankAccounts {
			if plaidBankAccount.GetPlaidAccountId() == bankAccount.PlaidAccountId {
				bankAccountIds = append(bankAccountIds, bankAccount.PlaidAccountId)
				plaidIdsToBankIds[bankAccount.PlaidAccountId] = bankAccount.Id
				plaidBankToLocalBank[bankAccount.PlaidAccountId] = bankAccount
				break
			}
		}

		// If an account is no longer visible in plaid that means that we won't receive updates for that account anymore. If
		// this happens, log something and mark that account as inactive. This way we can inform the user that the account
		// is no longer receiving updates.
		if _, ok := plaidIdsToBankIds[bankAccount.PlaidAccountId]; !ok {
			th.logger.Info("found bank account that is no longer present in plaid, it will be updated as inactive", zap.String("bankAccountId", bankAccount.PlaidAccountId))
			bankAccount.Status = apiv1.BankAccountStatus_BANK_ACCOUNT_STATUS_INACTIVE
			if err := postgresClient.UpdateBankAccount(ctx, bankAccount); err != nil {
				th.logger.Error("failed to update bank account", zap.Error(err))
			}
		}
	}

	if len(bankAccountIds) == 0 {
		th.logger.Warn("none of the linked bank accounts are active at plaid")
		return nil
	}

	plaidTransactions, err := plaidClient.GetAllTransactions(ctx, accessToken, startTime, endTime, bankAccountIds)
	if err = errors.Wrap(err, "failed to retrieve transactions from plaid for sync"); err != nil {
		th.logger.Error("failed to retrieve transactions from plaid for sync", zap.Error(err))
		return err
	}

	if len(plaidTransactions) == 0 {
		th.logger.Warn("no transactions were retrieved from plaid")
	}

	plaidTransactionIds := make([]string, 0, len(plaidTransactions))
	for _, transaction := range plaidTransactions {
		plaidTransactionIds = append(plaidTransactionIds, transaction.GetTransactionId())
	}

	// get all the transactions from clickhouse by transaction id
	clickhouseTxnIdToTxnMap := make(map[string]*apiv1.Transaction, 0)
	clickhousePlaidTransactions, err := clickhouseClient.GetTransactionsByPlaidTransactionIds(ctx, plaidTransactionIds)
	if err != nil {
		return err
	}

	for _, txn := range clickhousePlaidTransactions {
		clickhouseTxnIdToTxnMap[txn.GetTransactionId()] = txn
	}

	transactionsToUpdate := make([]*apiv1.Transaction, 0)
	transactionsToInsert := make([]*apiv1.Transaction, 0)
	for _, plaidTransaction := range plaidTransactions {
		amount := plaidTransaction.GetAmount()

		transactionName := plaidTransaction.GetName()

		// We only want to make the transaction name be the merchant name if the merchant name is shorter. This is
		// due to something I observed with a dominos transaction, where the merchant was improperly parsed and the
		// transaction ended up being called `Mnuslindstrom` rather than `Domino's`. This should fix that problem.
		if plaidTransaction.GetMerchantName() != "" && len(plaidTransaction.GetMerchantName()) < len(transactionName) {
			transactionName = plaidTransaction.GetMerchantName()
		}

		existingTransaction, ok := clickhouseTxnIdToTxnMap[plaidTransaction.GetTransactionId()]
		if !ok {
			txnRecord := &apiv1.Transaction{
				AccountId:                       plaidTransaction.GetAccountId(),
				Amount:                          amount,
				IsoCurrencyCode:                 plaidTransaction.GetIsoCurrencyCode(),
				UnofficialCurrencyCode:          plaidTransaction.GetUnofficialCurrencyCode(),
				CategoryId:                      plaidTransaction.GetCategoryId(),
				CheckNumber:                     plaidTransaction.GetCheckNumber(),
				Date:                            plaidTransaction.GetDate(),
				Datetime:                        plaidTransaction.GetDatetime(),
				AuthorizedDate:                  plaidTransaction.GetAuthorizedDate(),
				AuthorizedDatetime:              plaidTransaction.GetAuthorizedDatetime(),
				Name:                            transactionName,
				MerchantName:                    plaidTransaction.GetMerchantName(),
				PaymentChannel:                  plaidTransaction.GetPaymentChannel(),
				Pending:                         plaidTransaction.GetPending(),
				PendingTransactionId:            plaidTransaction.GetPendingTransactionId(),
				AccountOwner:                    plaidTransaction.GetAccountOwner(),
				TransactionId:                   plaidTransaction.GetTransactionId(),
				TransactionCode:                 plaidTransaction.GetTransactionCode(),
				Id:                              userId,
				UserId:                          payload.UserId,
				LinkId:                          link.Id,
				Sign:                            0,
				PersonalFinanceCategoryPrimary:  plaidTransaction.GetPersonalFinanceCategoryPrimary(),
				PersonalFinanceCategoryDetailed: plaidTransaction.GetPersonalFinanceCategoryDetailed(),
				LocationAddress:                 plaidTransaction.GetLocationAddress(),
				LocationCity:                    plaidTransaction.GetLocationCity(),
				LocationRegion:                  plaidTransaction.GetLocationRegion(),
				LocationPostalCode:              plaidTransaction.GetLocationPostalCode(),
				LocationCountry:                 plaidTransaction.GetLocationCountry(),
				LocationLat:                     plaidTransaction.GetLocationLat(),
				LocationLon:                     plaidTransaction.GetLocationLon(),
				LocationStoreNumber:             plaidTransaction.GetLocationStoreNumber(),
				PaymentMetaByOrderOf:            plaidTransaction.GetPaymentMetaByOrderOf(),
				PaymentMetaPayee:                plaidTransaction.GetPaymentMetaPayee(),
				PaymentMetaPayer:                plaidTransaction.GetPaymentMetaPayer(),
				PaymentMetaPaymentMethod:        plaidTransaction.GetPaymentMetaPaymentMethod(),
				PaymentMetaPaymentProcessor:     plaidTransaction.GetPaymentMetaPaymentProcessor(),
				PaymentMetaPpdId:                plaidTransaction.GetPaymentMetaPpdId(),
				PaymentMetaReason:               plaidTransaction.GetPaymentMetaReason(),
				PaymentMetaReferenceNumber:      plaidTransaction.GetPaymentMetaReferenceNumber(),
			}
			transactionsToInsert = append(transactionsToInsert, txnRecord)
			continue
		}

		var shouldUpdate bool
		if existingTransaction.Amount != amount {
			shouldUpdate = true
		}

		if existingTransaction.Pending != plaidTransaction.GetPending() {
			shouldUpdate = true
		}

		if strings.EqualFold(existingTransaction.PendingTransactionId, plaidTransaction.GetPendingTransactionId()) {
			shouldUpdate = true
		}

		existingTransaction.Amount = amount
		existingTransaction.Pending = plaidTransaction.GetPending()
		existingTransaction.PendingTransactionId = plaidTransaction.GetPendingTransactionId()

		if shouldUpdate {
			transactionsToUpdate = append(transactionsToUpdate, existingTransaction)
		}
	}

	if len(transactionsToUpdate) > 0 {
		th.logger.Info("updating transactions", zap.Int("count", len(transactionsToUpdate)))
		if err = clickhouseClient.UpdateTransactions(ctx, &payload.UserId, transactionsToUpdate); err != nil {
			return err
		}
	}

	if len(transactionsToInsert) > 0 {
		th.logger.Info("inserting transactions", zap.Int("count", len(transactionsToInsert)))
		// Reverse the list so the oldest records are inserted first.
		for i, j := 0, len(transactionsToInsert)-1; i < j; i, j = i+1, j-1 {
			transactionsToInsert[i], transactionsToInsert[j] = transactionsToInsert[j], transactionsToInsert[i]
		}
		if err := clickhouseClient.AddTransactions(ctx, &payload.UserId, transactionsToInsert); err != nil {
			return err
		}
	}

	if len(transactionsToInsert)+len(transactionsToUpdate) > 0 {
		updatedBankAccounts := make([]*apiv1.BankAccount, 0, len(plaidBankAccounts))
		for _, singleBankAccount := range plaidBankAccounts {
			bankAccount, ok := plaidBankToLocalBank[singleBankAccount.GetPlaidAccountId()]
			if !ok {
				th.logger.Warn("bank was not found in map", zap.Any("plaidBankAccountId", singleBankAccount.GetId()))
				continue
			}

			shouldUpdate := false
			available := singleBankAccount.GetBalance()
			current := singleBankAccount.GetCurrentFunds()

			if bankAccount.CurrentFunds != current {
				shouldUpdate = true
			}

			if bankAccount.Balance != available {
				shouldUpdate = true
			}

			plaidName := bankAccount.Name
			if bankAccount.Name != singleBankAccount.GetName() {
				plaidName = singleBankAccount.GetName()
				shouldUpdate = true
			}

			if shouldUpdate {
				bankAccount.Name = plaidName
				bankAccount.Balance = available
				bankAccount.CurrentFunds = current
				updatedBankAccounts = append(updatedBankAccounts, bankAccount)
			}
		}

		if err = postgresClient.UpdateBankAccounts(ctx, updatedBankAccounts); err != nil {
			th.logger.Error("failed to update bank accounts", zap.Error(err))
		}
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
