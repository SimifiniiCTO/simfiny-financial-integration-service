package taskhandler

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/SimifiniiCTO/asynq"
	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"go.uber.org/zap"
)

func NewSyncRecurringTransactionsTask(userId uint64, accessToken string, linkId uint64) (*asynq.Task, error) {
	payload, err := json.Marshal(&SyncPlaidTaskPayload{
		UserId:      userId,
		AccessToken: accessToken,
		LinkId:      linkId,
	})

	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TaskSyncRecurringTransactions.String(), payload), nil
}

func (th *TaskHandler) RunSyncRecurringTransactionsTask(ctx context.Context, task *asynq.Task) error {
	var (
		payload SyncPlaidTaskPayload
	)

	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return err
	}

	// sync the recurring transactions for the given user
	if err := th.pullRecurringTransactions(ctx, &payload.UserId, &payload.LinkId, &payload.AccessToken); err != nil {
		th.logger.Error("failed to pull recurring transactions", zap.Error(err))
		return err
	}

	return nil
}

func (th *TaskHandler) pullRecurringTransactions(ctx context.Context, userId, linkId *uint64, accessToken *string) error {
	var (
		plaidClient = th.plaidClient
	)
	// sync the recurring transactions for the given user
	newRecurringTransactions, err :=
		plaidClient.
			GetRecurringTransactions(
				ctx,
				accessToken,
				userId,
				linkId)
	if err != nil {
		th.logger.Error("failed to get recurring transactions", zap.Error(err))
		return err
	}

	// get all the old recurring transactions
	// create a hash map of the old recurring transactions
	// compare the old recurring transactions with the new ones
	// check if the new recurring transaction is in the old recurring transactions
	// if it is, then we need to update it
	// check if the new recurring transaction is different from the old one
	// The line `newRecurringTransaction.LinkId = oldRecurringTransaction` is assigning the value of
	// `oldRecurringTransaction` to the `LinkId` field of `newRecurringTransaction`. This is done to
	// ensure that the `LinkId` value is preserved when updating the recurring transaction.
	// if the new recurring transaction is not in the old recurring transactions, then we need to add it
	// store the recurring transactions in the database
	// TODO: need to condition on which recurring transactions are new and which are updates as well as which are to be deleted
	if err := th.pullRecurringTransactionsHelper(ctx, userId, newRecurringTransactions); err != nil {
		th.logger.Error("failed to pull recurring transactions", zap.Error(err))
		return err
	}

	return nil
}

func (th *TaskHandler) pullRecurringTransactionsHelper(ctx context.Context, userId *uint64, newRecurringTransactions []*schema.ReOccuringTransaction) error {
	var (
		clickhouseClient = th.clickhouseDb
	)

	recurringTransactionsInDatabase, err := clickhouseClient.GetUserReOccurringTransactions(ctx, userId)
	if err != nil {
		th.logger.Error("failed to get old recurring transactions", zap.Error(err))
		return err
	}

	txnF := func(subscription *schema.ReOccuringTransaction) string {
		merchantName := strings.ToLower(subscription.MerchantName)
		return fmt.Sprintf("%s-%s-%s-%s", subscription.AccountId, subscription.StreamId, subscription.CategoryId, merchantName)
	}

	recurringTransactionsInDatabaseMap := make(map[string]*schema.ReOccuringTransaction)
	for _, recurringTransaction := range recurringTransactionsInDatabase {
		key := txnF(recurringTransaction)
		recurringTransactionsInDatabaseMap[key] = recurringTransaction
	}

	subscriptionsToAdd := make([]*schema.ReOccuringTransaction, 0)
	subscriptionsToUpdate := make([]*schema.ReOccuringTransaction, 0)

	for _, newRecurringTransaction := range newRecurringTransactions {

		key := txnF(newRecurringTransaction)
		if oldRecurringTransaction, ok := recurringTransactionsInDatabaseMap[key]; ok {

			if newRecurringTransaction.StreamId != oldRecurringTransaction.StreamId ||
				newRecurringTransaction.CategoryId != oldRecurringTransaction.CategoryId ||
				newRecurringTransaction.Description != oldRecurringTransaction.Description ||
				newRecurringTransaction.MerchantName != oldRecurringTransaction.MerchantName ||
				newRecurringTransaction.LastAmount != oldRecurringTransaction.LastAmount ||
				newRecurringTransaction.LastDate != oldRecurringTransaction.LastDate ||
				newRecurringTransaction.PersonalFinanceCategoryPrimary != oldRecurringTransaction.PersonalFinanceCategoryPrimary {

				newRecurringTransaction.LinkId = oldRecurringTransaction.LinkId
				newRecurringTransaction.UserId = oldRecurringTransaction.UserId
				newRecurringTransaction.Flow = oldRecurringTransaction.Flow
				newRecurringTransaction.Id = oldRecurringTransaction.Id
				newRecurringTransaction.Sign = oldRecurringTransaction.Sign
				newRecurringTransaction.Time = oldRecurringTransaction.Time
				newRecurringTransaction.AdditionalProperties = oldRecurringTransaction.AdditionalProperties

				subscriptionsToUpdate = append(subscriptionsToUpdate, newRecurringTransaction)
			}
		} else {

			subscriptionsToAdd = append(subscriptionsToAdd, newRecurringTransaction)
		}
	}

	if len(subscriptionsToAdd) > 0 {
		if err := clickhouseClient.AddReOccurringTransactions(ctx, userId, subscriptionsToAdd); err != nil {
			th.logger.Error("failed to add recurring transactions", zap.Error(err))
			return err
		}
	}

	if len(subscriptionsToUpdate) > 0 {
		if err := clickhouseClient.UpdateReOccurringTransactions(ctx, userId, subscriptionsToUpdate); err != nil {
			th.logger.Error("failed to update recurring transactions", zap.Error(err))
			return err
		}
	}
	return nil
}
