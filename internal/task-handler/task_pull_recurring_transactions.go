package taskhandler

import (
	"context"
	"encoding/json"

	"github.com/SimifiniiCTO/asynq"
	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type PullUpdatedRecurringTransactionsTaskPayload struct {
	UserId      uint64   `json:"user_id"`
	LinkId      uint64   `json:"link_id"`
	AccessToken string   `json:"access_token"`
	AccountIds  []string `json:"account_ids"`
}

func NewPullUpdatedReCurringTransactionsTask(userId, linkId uint64, accessToken string, accountIds []string) (*asynq.Task, error) {
	if userId == 0 {
		return nil, errors.New("invalid input argument. user id cannot be empty")
	}

	if linkId == 0 {
		return nil, errors.New("invalid input argument. link id cannot be empty")
	}

	if accessToken == "" {
		return nil, errors.New("invalid input argument. access token cannot be empty")
	}

	rec := &PullUpdatedRecurringTransactionsTaskPayload{
		UserId:      userId,
		LinkId:      linkId,
		AccessToken: accessToken,
		AccountIds:  accountIds,
	}

	payload, err := json.Marshal(rec)
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TaskPullUpdatedRecurringTransactions.String(), payload), nil
}

func (th *TaskHandler) RunPullUpdatedReCurringTransactionsTask(ctx context.Context, t *asynq.Task) error {
	var (
		payload        = &PullUpdatedRecurringTransactionsTaskPayload{}
		postgresClient = th.postgresDb
		plaidClient    = th.plaidClient
	)

	if err := json.Unmarshal(t.Payload(), payload); err != nil {
		return err
	}

	userId := payload.UserId
	linkId := payload.LinkId
	accessToken := payload.AccessToken
	accountIds := payload.AccountIds

	// query the link from the database
	link, err := postgresClient.GetLink(ctx, userId, linkId, false)
	if err != nil {
		return err
	}

	if link.PlaidLink == nil {
		th.logger.Warn("provided link does not have any plaid credentials")
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

	if err := th.queryAndStoreRecurringTransactions(ctx, accessToken, userId, linkId, accountIds); err != nil {
		return err
	}

	return nil
}

// TODO: Rework Query Logic
func (th *TaskHandler) queryAndStoreRecurringTransactions(ctx context.Context, accessToken string, userId uint64, linkId uint64, accountIds []string) error {
	// query the re-curring transactions for the given accountId
	// get the existing re-occurring transactions from the database
	// updated set of transactions
	// new set of transactions
	// deleted set of transactions
	// copy the existing transaction id to the new transaction
	// delete the re-occurring transactions from the database
	// add the new re-occurring transactions to the database
	// update the re-occurring transactions in the database
	// add the transactions to the database
	var (
		plaidClient      = th.plaidClient
		clickhouseClient = th.clickhouseDb
	)
	transactions, err :=
		plaidClient.
			GetRecurringTransactionsForAccounts(ctx, &accessToken, &userId, &linkId, accountIds)
	if err != nil {
		th.logger.Error("error while getting re-occurring transactions from plaid", zap.Error(err))
		return err
	}

	currentTransactions, err := clickhouseClient.GetUserReOccurringTransactions(ctx, &userId)
	if err != nil {
		th.logger.Error("error while getting re-occurring transactions from database", zap.Error(err))
		return err
	}

	newTransactions := make([]*schema.ReOccuringTransaction, 0, len(transactions))
	for _, transaction := range transactions {
		if !isRecurringTransactionInSlice(transaction, currentTransactions) {
			newTransactions = append(newTransactions, transaction)
		}
	}

	updatedTransactions := make([]*schema.ReOccuringTransaction, 0, len(transactions))
	for _, transaction := range transactions {
		if isRecurringTransactionInSlice(transaction, currentTransactions) {
			transaction.UserId = userId
			transaction.LinkId = linkId
			transaction.Sign = 1
			updatedTransactions = append(updatedTransactions, transaction)
		}
	}

	deletedTransactionIds := make([]string, 0, len(currentTransactions))
	for _, transaction := range currentTransactions {
		if !isRecurringTransactionInSlice(transaction, transactions) {
			deletedTransactionIds = append(deletedTransactionIds, transaction.Id)
		}
	}

	if len(deletedTransactionIds) > 0 {
		th.logger.Info("deleting re-occurring transactions", zap.Any("deleted_recurring_txn_set", len(deletedTransactionIds)))
		if err := clickhouseClient.DeleteReOccurringTransactionsByIds(ctx, deletedTransactionIds); err != nil {
			return err
		}
	}

	if len(newTransactions) > 0 {
		th.logger.Info("adding new re-occurring transactions", zap.Any("new_recurring_txn_set", len(newTransactions)))
		if err := clickhouseClient.AddReOccurringTransactions(ctx, &userId, newTransactions); err != nil {
			return err
		}
	}

	if len(updatedTransactions) > 0 {
		th.logger.Info("updating re-occurring transactions", zap.Any("updated_recurring_txn_set", len(updatedTransactions)))
		if err := clickhouseClient.UpdateReOccurringTransactions(ctx, &userId, updatedTransactions); err != nil {
			return err
		}
	}

	return nil
}

func isRecurringTransactionInSlice(transaction *schema.ReOccuringTransaction, transactions []*schema.ReOccuringTransaction) bool {
	for _, t := range transactions {
		if t.CategoryId == transaction.CategoryId &&
			t.Description == transaction.Description &&
			t.MerchantName == transaction.MerchantName &&
			t.AverageAmount == transaction.AverageAmount &&
			t.LastAmount == transaction.LastAmount &&
			t.AccountId == transaction.AccountId {
			return true
		}
	}

	return false
}
