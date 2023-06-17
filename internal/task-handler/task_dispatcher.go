package taskhandler

import (
	"context"
	"time"

	"github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"
	taskprocessor "github.com/SimifiniiCTO/simfiny-core-lib/task-processor"
	"go.uber.org/zap"
)

// This is a method defined on a gRPC server struct that triggers a Plaid sync task for a given user.
// It takes in a context, user ID, Plaid link item ID, and access token as parameters, creates a new
// sync task using the taskhandler package, and enqueues the task using the taskprocessor. It also logs
// information about the enqueued task using a logger.
func DispatchPlaidSyncTask(ctx context.Context, tp *taskprocessor.TaskProcessor, instrumentation *instrumentation.Client, logger *zap.Logger, userId uint64, linkId uint64, accessToken string) error {
	if instrumentation != nil {
		txn := instrumentation.GetTraceFromContext(ctx)
		span := instrumentation.StartDatastoreSegment(txn, "trigger-plaid-sync")
		defer span.End()
	}

	// create a task to sync the plaid item
	task, err := NewSyncPlaidTask(userId, accessToken, linkId)
	if err != nil {
		return err
	}

	// enqueue the task
	taskInfo, err := tp.EnqueueTask(ctx, task)
	if err != nil {
		return err
	}

	logger.Info("enqueue task", zap.Any("task", taskInfo))
	return nil
}

func DispatchPullTransactionsTask(ctx context.Context, tp *taskprocessor.TaskProcessor, instrumentation *instrumentation.Client, logger *zap.Logger, userId uint64, linkId uint64, accessToken string, startTime, endTime time.Time) error {
	if instrumentation != nil {
		txn := instrumentation.GetTraceFromContext(ctx)
		span := instrumentation.StartDatastoreSegment(txn, "trigger-pull-transactions")
		defer span.End()
	}

	// create a task to sync the plaid item
	task, err := NewPullTransactionsTask(userId, linkId, accessToken, startTime, endTime)
	if err != nil {
		return err
	}

	// enqueue the task
	taskInfo, err := tp.EnqueueTask(ctx, task)
	if err != nil {
		return err
	}

	logger.Info("enqueue task", zap.Any("task", taskInfo))

	return nil
}

func DispatchPullUpdatedReCurringTransactionsTask(ctx context.Context, tp *taskprocessor.TaskProcessor, instrumentation *instrumentation.Client, logger *zap.Logger, userId, linkId uint64, accessToken string, accountIds []string) error {
	if instrumentation != nil {
		txn := instrumentation.GetTraceFromContext(ctx)
		span := instrumentation.StartDatastoreSegment(txn, "trigger-pull-recurring-transactions")
		defer span.End()
	}

	if len(accountIds) == 0 {
		logger.Error("no account ids found in webhook")
		return nil
	}

	task, err := NewPullUpdatedReCurringTransactionsTask(userId, linkId, accessToken, accountIds)
	if err != nil {
		return err
	}

	// enqueue the task
	taskInfo, err := tp.EnqueueTask(ctx, task)
	if err != nil {
		return err
	}

	logger.Info("enqueue task", zap.Any("task", taskInfo))

	return nil
}

func DispatchPullInvestmentTransactionsTask(ctx context.Context, tp *taskprocessor.TaskProcessor, instrumentation *instrumentation.Client, logger *zap.Logger, userId, linkId uint64, accessToken string, accountIds []string) error {
	if instrumentation != nil {
		txn := instrumentation.GetTraceFromContext(ctx)
		span := instrumentation.StartDatastoreSegment(txn, "trigger-pull-investment-transactions")
		defer span.End()
	}

	if len(accountIds) == 0 {
		logger.Error("no account ids found in webhook")
		return nil
	}

	task, err := NewPullInvestmentTransactionsTask(userId, linkId, accessToken, accountIds)
	if err != nil {
		return err
	}

	// enqueue the task
	taskInfo, err := tp.EnqueueTask(ctx, task)
	if err != nil {
		return err
	}

	logger.Info("enqueue task", zap.Any("task", taskInfo))

	return nil
}

func DispatchPullInvestmentHoldingsTask(ctx context.Context, tp *taskprocessor.TaskProcessor, instrumentation *instrumentation.Client, logger *zap.Logger, userId, linkId uint64, accessToken string, accountIds []string) error {
	if instrumentation != nil {
		txn := instrumentation.GetTraceFromContext(ctx)
		span := instrumentation.StartDatastoreSegment(txn, "trigger-pull-investment-holdings")
		defer span.End()
	}

	if len(accountIds) == 0 {
		logger.Error("no account ids found in webhook")
		return nil
	}

	task, err := NewPullInvestmentHoldingsTask(userId, linkId, accessToken, accountIds)
	if err != nil {
		return err
	}

	// enqueue the task
	taskInfo, err := tp.EnqueueTask(ctx, task)
	if err != nil {
		return err
	}

	logger.Info("enqueue task", zap.Any("task", taskInfo))

	return nil
}

func DispatchSyncLiabilityAccountsTask(ctx context.Context, tp *taskprocessor.TaskProcessor, instrumentation *instrumentation.Client, logger *zap.Logger, userId, linkId uint64, accessToken string, accountIds []string) error {
	if instrumentation != nil {
		txn := instrumentation.GetTraceFromContext(ctx)
		span := instrumentation.StartDatastoreSegment(txn, "trigger-sync-new-liaibility-accounts")
		defer span.End()
	}

	if len(accountIds) == 0 {
		logger.Error("no account ids found in webhook")
		return nil
	}

	task, err := NewSyncNewLiabilityAccountsTask(userId, accessToken, linkId, accountIds)
	if err != nil {
		return err
	}

	// enqueue the task
	taskInfo, err := tp.EnqueueTask(ctx, task)
	if err != nil {
		return err
	}

	logger.Info("enqueue task", zap.Any("task", taskInfo))

	return nil
}
