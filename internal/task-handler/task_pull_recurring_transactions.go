package taskhandler

import (
	"context"
	"encoding/json"

	"github.com/SimifiniiCTO/asynq"
	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/pkg/errors"
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
		payload          = &PullUpdatedRecurringTransactionsTaskPayload{}
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

	// query the re-curring transactions for the given accountId
	newOrUpdatedReCurringTxns, err := plaidClient.GetRecurringTransactionsForAccounts(ctx, &accessToken, &userId, &linkId, accountIds)
	if err != nil {
		return err
	}

	// get the existing re-occurring transactions from the database
	existingReCurringTxns, err := clickhouseClient.GetUserReOccurringTransactions(ctx, &userId)
	if err != nil {
		return err
	}

	// updated set of transactions
	updatedRecurringTxnSet := make([]*schema.ReOccuringTransaction, len(existingReCurringTxns))
	// new set of transactions
	newRecurringTxnSet := make([]*schema.ReOccuringTransaction, len(existingReCurringTxns))
	// deleted set of transactions
	deletedRecurringTxnSet := make([]uint64, len(existingReCurringTxns))

	for _, newTxn := range newOrUpdatedReCurringTxns {
		for idx, existingTxn := range existingReCurringTxns {
			if newTxn.MerchantName == existingTxn.MerchantName &&
				newTxn.Frequency == existingTxn.Frequency &&
				newTxn.AverageAmount == existingTxn.AverageAmount &&
				newTxn.PersonalFinanceCategoryDetailed == existingTxn.PersonalFinanceCategoryDetailed {
				// copy the existing transaction id to the new transaction
				updatedRecurringTxnSet = append(updatedRecurringTxnSet, copyOldtxnToNewtxn(existingTxn, newTxn))
				break
			} else if idx == len(existingReCurringTxns)-1 {
				newTxn.UserId = userId
				newTxn.LinkId = linkId
				newRecurringTxnSet = append(newRecurringTxnSet, newTxn)
			}
		}
	}

	for _, existingTxn := range existingReCurringTxns {
		for idx, newTxn := range newOrUpdatedReCurringTxns {
			if idx == len(newOrUpdatedReCurringTxns)-1 && newTxn.MerchantName != existingTxn.MerchantName &&
				newTxn.Frequency != existingTxn.Frequency &&
				newTxn.AverageAmount != existingTxn.AverageAmount &&
				newTxn.PersonalFinanceCategoryDetailed != existingTxn.PersonalFinanceCategoryDetailed {
				deletedRecurringTxnSet = append(deletedRecurringTxnSet, existingTxn.Id)
			}
		}
	}

	if len(deletedRecurringTxnSet) > 0 {
		// delete the re-occurring transactions from the database
		if err := clickhouseClient.DeleteReOccurringTransactionsByIds(ctx, deletedRecurringTxnSet); err != nil {
			return err
		}
	}

	// add the new re-occurring transactions to the database
	if err := clickhouseClient.AddReOccurringTransactions(ctx, &userId, newRecurringTxnSet); err != nil {
		return err
	}

	// update the re-occurring transactions in the database
	// add the transactions to the database
	if err := clickhouseClient.UpdateReOccurringTransactions(ctx, &userId, updatedRecurringTxnSet); err != nil {
		return err
	}

	return nil
}

func copyOldtxnToNewtxn(oldtxn *schema.ReOccuringTransaction, newtxn *schema.ReOccuringTransaction) *schema.ReOccuringTransaction {
	return &schema.ReOccuringTransaction{
		AccountId:                       newtxn.AccountId,
		StreamId:                        newtxn.StreamId,
		CategoryId:                      newtxn.CategoryId,
		Description:                     newtxn.CategoryId,
		MerchantName:                    newtxn.MerchantName,
		PersonalFinanceCategoryPrimary:  newtxn.PersonalFinanceCategoryPrimary,
		PersonalFinanceCategoryDetailed: newtxn.PersonalFinanceCategoryDetailed,
		FirstDate:                       newtxn.FirstDate,
		LastDate:                        newtxn.LastDate,
		Frequency:                       newtxn.Frequency,
		TransactionIds:                  newtxn.TransactionIds,
		AverageAmount:                   newtxn.AverageAmount,
		AverageAmountIsoCurrencyCode:    newtxn.AverageAmountIsoCurrencyCode,
		LastAmount:                      newtxn.LastAmount,
		LastAmountIsoCurrencyCode:       newtxn.LastAmountIsoCurrencyCode,
		IsActive:                        newtxn.IsActive,
		Status:                          newtxn.Status,
		UpdatedTime:                     newtxn.UpdatedTime,
		Flow:                            newtxn.Flow,
		UserId:                          oldtxn.UserId,
		LinkId:                          oldtxn.LinkId,
		Id:                              oldtxn.Id,
	}
}
