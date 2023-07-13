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

func (th *TaskHandler) processSyncOperation(ctx context.Context, userId, linkId uint64, accessToken, trigger string) ([]*apiv1.BankAccount, error) {
	var (
		postgresClient   = th.postgresDb
		clickhouseClient = th.clickhouseDb
		plaidClient      = th.plaidClient
	)

	// query the link from the database
	link, err := postgresClient.GetLink(ctx, userId, linkId, false)
	if err != nil {
		return nil, err
	}

	if link.PlaidLink == nil {
		th.logger.Warn("provided link does not have any plaid credentials")
		return nil, errors.New("plaid link not found")
	}

	// query plaid for all of the user's accounts
	plaidBankAccounts, err := plaidClient.GetAccounts(ctx, accessToken, userId)
	if err != nil {
		return nil, err
	}

	// get the liability accounts
	// TODO: need to figure out wether to update the existing investment account or add a new one
	investmentAccounts, err := plaidClient.GetInvestmentAccount(ctx, userId, accessToken)
	if err != nil {
		return nil, err
	}

	if len(investmentAccounts) > 0 {
		th.logger.Info("found investment accounts", zap.Int("count", len(investmentAccounts)))
	}

	// get the credit accounts
	// TODO: need to figure out wether to update the existing credit account or add a new one
	creditAccountsSet, err := plaidClient.GetPlaidLiabilityAccounts(ctx, &accessToken)
	if err != nil {
		return nil, err
	}

	if len(creditAccountsSet.CrediCardAccounts) > 0 {
		syncedCreditAccounts := creditAccountsSet.CrediCardAccounts
		currentCreditAccounts := link.GetCreditAccounts()

		// iterate over the syncred credit accounts and cross reference them with the existing credit accounts
		// if the synced credit account is not found in the existing credit accounts, then add it to the existing
		// credit accounts
		creditCardAccountsToUpdate := make([]*apiv1.CreditAccountORM, 0, len(syncedCreditAccounts))
		creditCardAccountsToAdd := make([]*apiv1.CreditAccountORM, 0, len(syncedCreditAccounts))
		for _, syncedCreditAccount := range syncedCreditAccounts {
			found := false
			for idx, currentCreditAccount := range currentCreditAccounts {
				if syncedCreditAccount.PlaidAccountId == currentCreditAccount.PlaidAccountId {
					found = true
					break
				} else if syncedCreditAccount.PlaidAccountId != currentCreditAccount.PlaidAccountId && idx == len(currentCreditAccounts)-1 {
					acctOrm, err := syncedCreditAccount.ToORM(ctx)
					if err != nil {
						return nil, err
					}

					creditCardAccountsToAdd = append(creditCardAccountsToAdd, &acctOrm)
				}
			}

			if found {
				acctOrm, err := syncedCreditAccount.ToORM(ctx)
				if err != nil {
					return nil, err
				}
				creditCardAccountsToUpdate = append(creditCardAccountsToUpdate, &acctOrm)
			}
		}

		l := th.postgresDb.QueryOperator.LinkORM
		linkOrm, err := link.ToORM(ctx)
		if err != nil {
			return nil, err
		}

		// update the credit accounts in the database
		if err := l.CreditAccounts.Model(&linkOrm).Replace(creditCardAccountsToUpdate...); err != nil {
			return nil, err
		}

		// add the new credit accounts to the database
		if err := l.CreditAccounts.Model(&linkOrm).Append(creditCardAccountsToAdd...); err != nil {
			return nil, err
		}

		th.logger.Info("found credit accounts", zap.Int("count", len(creditAccountsSet.CrediCardAccounts)))
	}

	if len(creditAccountsSet.MortgageLoanAccts) > 0 {
		syncedMortgageLoanAccounts := creditAccountsSet.MortgageLoanAccts
		currentMortgageLoanAccounts := link.GetMortgageAccounts()

		// iterate over the syncred credit accounts and cross reference them with the existing credit accounts
		// if the synced credit account is not found in the existing credit accounts, then add it to the existing
		// credit accounts
		mortgageLoanAccountsToUpdate := make([]*apiv1.MortgageAccountORM, 0, len(syncedMortgageLoanAccounts))
		mortgageLoanAccountsToAdd := make([]*apiv1.MortgageAccountORM, 0, len(syncedMortgageLoanAccounts))
		for _, syncedCreditAccount := range syncedMortgageLoanAccounts {
			found := false
			for idx, currentCreditAccount := range currentMortgageLoanAccounts {
				if syncedCreditAccount.PlaidAccountId == currentCreditAccount.PlaidAccountId {
					found = true
					break
				} else if syncedCreditAccount.PlaidAccountId != currentCreditAccount.PlaidAccountId && idx == len(currentMortgageLoanAccounts)-1 {
					acctOrm, err := syncedCreditAccount.ToORM(ctx)
					if err != nil {
						return nil, err
					}

					mortgageLoanAccountsToAdd = append(mortgageLoanAccountsToAdd, &acctOrm)
				}
			}

			if found {
				acctOrm, err := syncedCreditAccount.ToORM(ctx)
				if err != nil {
					return nil, err
				}
				mortgageLoanAccountsToUpdate = append(mortgageLoanAccountsToUpdate, &acctOrm)
			}
		}

		l := th.postgresDb.QueryOperator.LinkORM
		linkOrm, err := link.ToORM(ctx)
		if err != nil {
			return nil, err
		}

		// update the credit accounts in the database
		if err := l.MortgageAccounts.Model(&linkOrm).Replace(mortgageLoanAccountsToUpdate...); err != nil {
			return nil, err
		}

		// add the new credit accounts to the database
		if err := l.MortgageAccounts.Model(&linkOrm).Append(mortgageLoanAccountsToAdd...); err != nil {
			return nil, err
		}

		th.logger.Info("found mortgage accounts", zap.Int("count", len(creditAccountsSet.MortgageLoanAccts)))
	}

	if len(creditAccountsSet.StudentLoanAccts) > 0 {
		syncedStudentLoanAccounts := creditAccountsSet.StudentLoanAccts
		currentStudentLoanAccounts := link.GetStudentLoanAccounts()

		// iterate over the syncred credit accounts and cross reference them with the existing credit accounts
		// if the synced credit account is not found in the existing credit accounts, then add it to the existing
		// credit accounts
		studentLoanAccountsToUpdate := make([]*apiv1.StudentLoanAccountORM, 0, len(syncedStudentLoanAccounts))
		studentLoanAccountsToAdd := make([]*apiv1.StudentLoanAccountORM, 0, len(syncedStudentLoanAccounts))
		for _, syncedCreditAccount := range syncedStudentLoanAccounts {
			found := false
			for idx, currentCreditAccount := range currentStudentLoanAccounts {
				if syncedCreditAccount.PlaidAccountId == currentCreditAccount.PlaidAccountId {
					found = true
					break
				} else if syncedCreditAccount.PlaidAccountId != currentCreditAccount.PlaidAccountId && idx == len(currentStudentLoanAccounts)-1 {
					acctOrm, err := syncedCreditAccount.ToORM(ctx)
					if err != nil {
						return nil, err
					}

					studentLoanAccountsToAdd = append(studentLoanAccountsToAdd, &acctOrm)
				}
			}

			if found {
				acctOrm, err := syncedCreditAccount.ToORM(ctx)
				if err != nil {
					return nil, err
				}
				studentLoanAccountsToUpdate = append(studentLoanAccountsToUpdate, &acctOrm)
			}
		}

		l := th.postgresDb.QueryOperator.LinkORM
		linkOrm, err := link.ToORM(ctx)
		if err != nil {
			return nil, err
		}

		// update the credit accounts in the database
		if err := l.StudentLoanAccounts.Model(&linkOrm).Replace(studentLoanAccountsToUpdate...); err != nil {
			return nil, err
		}

		// add the new credit accounts to the database
		if err := l.StudentLoanAccounts.Model(&linkOrm).Append(studentLoanAccountsToAdd...); err != nil {
			return nil, err
		}

		th.logger.Info("found student loan accounts", zap.Int("count", len(creditAccountsSet.StudentLoanAccts)))
	}

	if len(plaidBankAccounts) == 0 {
		th.logger.Warn("no accounts found for user", zap.Uint64("user_id", userId))
		return nil, nil
	}

	plaidIDtoBankIDMap := make(map[string]uint64, 0)
	plaidBankIDToBankRecordMap := make(map[uint64]*apiv1.BankAccount, 0)
	bankAccountIds := make([]uint64, 0, len(plaidBankAccounts))

	for _, customBankAccount := range link.GetBankAccounts() {
		for _, plaidBankAccount := range plaidBankAccounts {
			if plaidBankAccount.PlaidAccountId == customBankAccount.PlaidAccountId {
				bankAccountIds = append(bankAccountIds, customBankAccount.Id)
				plaidIDtoBankIDMap[plaidBankAccount.PlaidAccountId] = customBankAccount.Id
				plaidBankIDToBankRecordMap[customBankAccount.Id] = customBankAccount
				break
			}
		}

		// If an account is no longer visible in plaid that means that we won't receive updates for that account anymore. If
		// this happens, log something and mark that account as inactive. This way we can inform the user that the account
		// is no longer receiving updates.
		if _, ok := plaidIDtoBankIDMap[customBankAccount.PlaidAccountId]; !ok {
			customBankAccount.Status = apiv1.BankAccountStatus_BANK_ACCOUNT_STATUS_INACTIVE
			// update the bank account in the database
			// TODO: maybe emit a metric here
			if err := postgresClient.UpdateBankAccount(ctx, customBankAccount); err != nil {
				return nil, err
			}
		}
	}

	if (len(bankAccountIds)) == 0 {
		th.logger.Warn("none of the linked bank accounts are active at plaid", zap.Uint64("user_id", userId))
		return nil, nil
	}

	var cursor *string
	// get the last plaid sync
	lasPlaidSync := link.GetPlaidSync()
	if lasPlaidSync != nil {
		cursor = &lasPlaidSync.NextCursor
	}

	for iter := 0; iter < 10; iter++ {
		syncResult, err := plaidClient.Sync(ctx, cursor, &accessToken)
		if err != nil {
			return nil, errors.Wrap(err, "failed to sync with plaid")
		}

		th.logger.Info("synced result", zap.Any("result", syncResult))

		// record the sync event in the database
		linkId := link.Id
		nextCursor := syncResult.NextCursor
		added := len(syncResult.New)
		deleted := len(syncResult.Deleted)
		modified := len(syncResult.Updated)
		if err := postgresClient.RecordPlaidSync(ctx, userId, linkId, trigger, nextCursor, int64(added), int64(modified), int64(deleted)); err != nil {
			return nil, err
		}

		// Update the cursor incase we need to iterate again.
		cursor = &syncResult.NextCursor

		// If we received nothing to insert/update/remove then do nothing
		if len(syncResult.New)+len(syncResult.Updated)+len(syncResult.Deleted) == 0 {
			log.Info("no new data from plaid, nothing to be done")
			return nil, nil
		}

		plaidTransactions := append(syncResult.New, syncResult.Updated...)
		th.logger.Info("received new transactions from plaid", zap.Int("count", len(plaidTransactions)))

		plaidTransactionIds := make([]string, 0, len(plaidTransactions))
		for _, transaction := range plaidTransactions {
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

			existingTransaction, ok := plaidTxnIdToTxnMap[plaidTransaction.GetTransactionId()]
			if !ok {
				txnRecord := &apiv1.Transaction{
					AccountId:                       plaidTransaction.GetAccountId(),
					Amount:                          amount,
					IsoCurrencyCode:                 plaidTransaction.GetIsoCurrencyCode(),
					UnofficialCurrencyCode:          plaidTransaction.GetUnofficialCurrencyCode(),
					CategoryId:                      plaidTransaction.GetCategoryId(),
					CheckNumber:                     plaidTransaction.GetCheckNumber(),
					CurrentDate:                     plaidTransaction.GetCurrentDate(),
					CurrentDatetime:                 plaidTransaction.GetCurrentDatetime(),
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
					Id:                              "",
					UserId:                          userId,
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
			if err = clickhouseClient.UpdateTransactions(ctx, &userId, transactionsToUpdate); err != nil {
				return nil, err
			}
		}

		if len(transactionsToInsert) > 0 {
			th.logger.Info("inserting transactions", zap.Int("count", len(transactionsToInsert)))
			if err := clickhouseClient.AddTransactions(ctx, &userId, transactionsToInsert); err != nil {
				return nil, err
			}
		}

		if len(transactionsToInsert)+len(transactionsToUpdate) > 0 {
			updatedBankAccounts := make([]*apiv1.BankAccount, 0, len(plaidBankAccounts))
			for _, singleBankAccount := range plaidBankAccounts {
				bankAccount, ok := plaidBankIDToBankRecordMap[singleBankAccount.GetId()]
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

		if !syncResult.HasMore {
			break
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
		return nil, err
	}

	if linkWasSetup {
		// TODO: Send the notification that the link has been set up.
		th.logger.Info("sending link setup notification")
	}

	return plaidBankAccounts, nil
}
