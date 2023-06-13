package grpc

import (
	"context"
	"time"

	taskhandler "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/task-handler"
	"go.uber.org/zap"
)

// This is a method defined on a gRPC server struct that triggers a Plaid sync task for a given user.
// It takes in a context, user ID, Plaid link item ID, and access token as parameters, creates a new
// sync task using the taskhandler package, and enqueues the task using the taskprocessor. It also logs
// information about the enqueued task using a logger.
func (s *Server) DispatchPlaidSyncTask(ctx context.Context, userId uint64, linkItemId, accessToken string) error {
	var (
		tp = s.taskprocessor
	)

	if s.instrumentation != nil {
		txn := s.instrumentation.GetTraceFromContext(ctx)
		span := s.instrumentation.StartDatastoreSegment(txn, "trigger-plaid-sync")
		defer span.End()
	}

	// create a task to sync the plaid item
	task, err := taskhandler.NewSyncPlaidTask(userId, linkItemId, accessToken)
	if err != nil {
		return err
	}

	// enqueue the task
	taskInfo, err := tp.EnqueueTask(ctx, task)
	if err != nil {
		return err
	}

	s.logger.Info("enqueue task", zap.Any("task", taskInfo))
	return nil
}

func (s *Server) DispatchPullTransactionsTask(ctx context.Context, userId uint64, linkId uint64, accessToken string, startTime, endTime time.Time) error {
	var (
		tp = s.taskprocessor
	)

	if s.instrumentation != nil {
		txn := s.instrumentation.GetTraceFromContext(ctx)
		span := s.instrumentation.StartDatastoreSegment(txn, "trigger-pull-transactions")
		defer span.End()
	}

	// create a task to sync the plaid item
	task, err := taskhandler.NewPullTransactionsTask(userId, linkId, accessToken, startTime, endTime)
	if err != nil {
		return err
	}

	// enqueue the task
	taskInfo, err := tp.EnqueueTask(ctx, task)
	if err != nil {
		return err
	}

	s.logger.Info("enqueue task", zap.Any("task", taskInfo))

	return nil
}

func (s *Server) DispatchPullUpdatedReCurringTransactionsTask(ctx context.Context, userId, linkId uint64, accessToken string, accountIds []string) error {
	var (
		tp = s.taskprocessor
	)

	if s.instrumentation != nil {
		txn := s.instrumentation.GetTraceFromContext(ctx)
		span := s.instrumentation.StartDatastoreSegment(txn, "trigger-pull-recurring-transactions")
		defer span.End()
	}

	if len(accountIds) == 0 {
		s.logger.Error("no account ids found in webhook")
		return nil
	}

	task, err := taskhandler.NewPullUpdatedReCurringTransactionsTask(userId, linkId, accessToken, accountIds)
	if err != nil {
		return err
	}

	// enqueue the task
	taskInfo, err := tp.EnqueueTask(ctx, task)
	if err != nil {
		return err
	}

	s.logger.Info("enqueue task", zap.Any("task", taskInfo))

	return nil
}
