package taskhandler

import (
	"context"
	"encoding/json"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"

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
		payload          = &PullInvestmentTransactionsTaskPayload{}
		postgresClient   = th.postgresDb
		clickhouseClient = th.clickhouseDb
		plaidClient      = th.plaidClient
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
	transactions, err := plaidClient.GetInvestmentTransactions(ctx, &accessToken, &userId, link, "", "", accountIds, 0, 0)
	if err != nil {
		return err
	}

	// need to cross reference the transactions obtained from plaid with the transactions in the database
	// to determine which transactions are new and which transactions are updates
	currentTransactions, err := clickhouseClient.GetAllInvestmentTransactions(ctx, &userId)
	if err != nil {
		return err
	}

	// get new transactions
	newTransactions := make([]*schema.InvestmentTransaction, 0, len(transactions))
	for _, transaction := range transactions {
		if !isTransactionInSlice(transaction, currentTransactions) {
			newTransactions = append(newTransactions, transaction)
		}
	}

	// get updated transactions
	updatedTransactions := make([]*schema.InvestmentTransaction, 0, len(transactions))
	for _, transaction := range transactions {
		if isTransactionInSlice(transaction, currentTransactions) {
			updatedTransactions = append(updatedTransactions, transaction)
		}
	}

	// get deleted transactions
	deletedTransactionIds := make([]uint64, 0, len(currentTransactions))
	for _, transaction := range currentTransactions {
		if !isTransactionInSlice(transaction, transactions) {
			deletedTransactionIds = append(deletedTransactionIds, transaction.Id)
		}
	}

	// save deleted transactions
	if err := clickhouseClient.DeleteInvestmentTransactions(ctx, deletedTransactionIds...); err != nil {
		return err
	}

	// save updated transactions
	if err := clickhouseClient.UpdateInvestmentTransactions(ctx, &userId, updatedTransactions); err != nil {
		return err
	}

	// save new transactions
	if err := clickhouseClient.AddInvestmentTransactions(ctx, &userId, newTransactions); err != nil {
		return err
	}

	// save investment transactions
	if err := clickhouseClient.AddInvestmentTransactions(ctx, &userId, transactions); err != nil {
		return err
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
