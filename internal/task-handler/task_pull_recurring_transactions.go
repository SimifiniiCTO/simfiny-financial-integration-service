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
	var (
		plaidClient = th.plaidClient
	)

	if len(accountIds) == 0 {
		return errors.New("invalid input argument. account ids cannot be empty")
	}

	transactions, err :=
		plaidClient.
			GetRecurringTransactionsForAccounts(ctx, &accessToken, &userId, &linkId, accountIds)
	if err != nil {
		th.logger.Error("error while getting re-occurring transactions from plaid", zap.Error(err))
		return err
	}

	if err := th.pullRecurringTransactionsHelper(ctx, &userId, transactions); err != nil {
		th.logger.Error("failed to pull recurring transactions", zap.Error(err))
		return err
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
