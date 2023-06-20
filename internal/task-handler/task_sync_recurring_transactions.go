package taskhandler

import (
	"context"
	"encoding/json"

	"github.com/SimifiniiCTO/asynq"
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
		clickhouseClient = th.clickhouseDb
		plaidClient      = th.plaidClient
	)
	// sync the recurring transactions for the given user
	recurringTransactions, err :=
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

	// store the recurring transactions in the database
	// TODO: need to condition on which recurring transactions are new and which are updates as well as which are to be deleted
	if err := clickhouseClient.AddReOccurringTransactions(ctx, userId, recurringTransactions); err != nil {
		th.logger.Error("failed to add recurring transactions", zap.Error(err))
		return err
	}

	return nil
}
