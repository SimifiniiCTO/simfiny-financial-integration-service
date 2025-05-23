package taskhandler

import (
	"context"
	"encoding/json"
	"time"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"go.uber.org/zap"

	"github.com/SimifiniiCTO/asynq"
	"github.com/pkg/errors"
)

type PullInvestmentTransactionsTaskPayload struct {
	UserId      uint64   `json:"user_id"`
	LinkId      uint64   `json:"link_id"`
	AccessToken string   `json:"access_token"`
	AccountIds  []string `json:"account_ids"`
}

func NewPullInvestmentTransactionsTask(userId, linkId uint64, accessToken string, accountIds []string) (*asynq.Task, error) {
	if userId == 0 {
		return nil, errors.New("invalid input argument. user id cannot be empty")
	}

	if linkId == 0 {
		return nil, errors.New("invalid input argument. link id cannot be empty")
	}

	if accessToken == "" {
		return nil, errors.New("invalid input argument. access token cannot be empty")
	}

	rec := &PullInvestmentTransactionsTaskPayload{
		UserId:      userId,
		LinkId:      linkId,
		AccessToken: accessToken,
		AccountIds:  accountIds,
	}

	payload, err := json.Marshal(rec)
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TaskPullInvestmentTransactions.String(), payload), nil
}

func (th *TaskHandler) RunPullInvestmentTransactionsTask(ctx context.Context, t *asynq.Task) error {
	var (
		payload        = &PullInvestmentTransactionsTaskPayload{}
		postgresClient = th.postgresDb
	)

	if err := json.Unmarshal(t.Payload(), payload); err != nil {
		return err
	}

	userId := payload.UserId
	linkId := payload.LinkId
	accessToken := payload.AccessToken
	accountIds := payload.AccountIds

	// get link
	link, err := postgresClient.GetLink(ctx, userId, linkId, false)
	if err != nil {
		return err
	}

	// get investment transactions
	// get last 2 weeks worth of transactions
	// TODO: track offset and number of transactions fetched in database
	// need to cross reference the transactions obtained from plaid with the transactions in the database
	// to determine which transactions are new and which transactions are updates
	// get new transactions
	// get updated transactions
	// get deleted transactions
	// save deleted transactions
	// save updated transactions
	// save new transactions
	if err := th.queryAndUpdateInvestmentTransactions(ctx, accessToken, userId, link, accountIds); err != nil {
		return err
	}

	return nil
}

func (th *TaskHandler) queryAndUpdateInvestmentTransactions(ctx context.Context, accessToken string, userId uint64, link *schema.Link, accountIds []string) error {
	// get investment transactions
	// get last 2 weeks worth of transactions
	// TODO: track offset and number of transactions fetched in database
	// need to cross reference the transactions obtained from plaid with the transactions in the database
	// to determine which transactions are new and which transactions are updates
	// get new transactions
	// get updated transactions
	// get deleted transactions
	// save deleted transactions
	// save updated transactions
	// save new transactions

	var (
		clickhouseClient = th.clickhouseDb
		plaidClient      = th.plaidClient
	)

	startTime := time.Now().AddDate(0, 0, -14).Format("2006-01-02")
	endTime := time.Now().Format("2006-01-02")

	th.logger.Info("getting investment transactions via plaid")
	transactions, err := plaidClient.GetInvestmentTransactions(ctx, &accessToken, &userId, link, startTime, endTime, accountIds, 0, 0)
	if err != nil {
		return err
	}

	th.logger.Info("got investment transactions via plaid", zap.Int("count", len(transactions)))
	currentTransactions, err := clickhouseClient.GetAllInvestmentTransactions(ctx, &userId)
	if err != nil {
		return err
	}

	th.logger.Info("got current investment transactions", zap.Int("count", len(currentTransactions)))
	newTransactions := make([]*schema.InvestmentTransaction, 0, len(transactions))
	for _, transaction := range transactions {
		if !isTransactionInSlice(transaction, currentTransactions) {
			newTransactions = append(newTransactions, transaction)
		}
	}

	updatedTransactions := make([]*schema.InvestmentTransaction, 0, len(transactions))
	for _, transaction := range transactions {
		if isTransactionInSlice(transaction, currentTransactions) {
			updatedTransactions = append(updatedTransactions, transaction)
		}
	}

	th.logger.Info("got updated investment transactions", zap.Int("count", len(updatedTransactions)))

	deletedTransactionIds := make([]string, 0, len(currentTransactions))
	for _, transaction := range currentTransactions {
		if !isTransactionInSlice(transaction, transactions) {
			deletedTransactionIds = append(deletedTransactionIds, transaction.Id)
		}
	}

	th.logger.Info("got deleted investment transactions", zap.Int("count", len(deletedTransactionIds)))

	if len(deletedTransactionIds) > 0 {
		if err := clickhouseClient.DeleteInvestmentTransactions(ctx, deletedTransactionIds...); err != nil {
			return err
		}
	}

	if len(updatedTransactions) > 0 {
		if err := clickhouseClient.UpdateInvestmentTransactions(ctx, &userId, updatedTransactions); err != nil {
			return err
		}
	}

	if len(newTransactions) > 0 {
		if err := clickhouseClient.AddInvestmentTransactions(ctx, &userId, newTransactions); err != nil {
			return err
		}
	}
	return nil
}

func isTransactionInSlice(transaction *schema.InvestmentTransaction, transactions []*schema.InvestmentTransaction) bool {
	for _, t := range transactions {
		if t.InvestmentTransactionId == transaction.InvestmentTransactionId {
			return true
		}
	}

	return false
}
