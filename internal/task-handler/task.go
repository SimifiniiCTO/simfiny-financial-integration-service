package taskhandler

import (
	"context"
	"fmt"

	"github.com/SimifiniiCTO/asynq"
	"github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"
	"github.com/SimifiniiCTO/simfiny-core-lib/task-processor/taskhandler"
	clickhousedb "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/clickhouse-database"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/plaidhandler"
	postgresdb "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/postgres-database"

	"go.uber.org/zap"
)

var (
	ErrInvalidPostgresConnectionObject   = fmt.Errorf("invalid postgres connection object")
	ErrInvalidClickhouseConnectionObject = fmt.Errorf("invalid clickhouse connection object")
	ErrInvalidInstrumentationClient      = fmt.Errorf("invalid instrumentation client")
	ErrInvalidTaskHandler                = fmt.Errorf("invalid task handler")
	ErrInvalidPlaidClient                = fmt.Errorf("invalid plaid client")
)

type TaskType string

const (
	TaskSyncTransactions                 TaskType = "transactions:sync"
	TaskDeleteLink                       TaskType = "link:delete"
	TaskDeleteTransactions               TaskType = "transactions:delete"
	TaskPullTransactions                 TaskType = "transactions:pull"
	TaskPullUpdatedRecurringTransactions TaskType = "transactions:pull:updated-recurring"
	TaskPullInvestmentTransactions       TaskType = "transactions:pull:investment"
	TaskPullInvestmentHoldings           TaskType = "holdings:pull:investment"
	TaskSyncNewLiabilityAccounts         TaskType = "accounts:sync:new-liability"
)

func (t TaskType) String() string {
	return string(t)
}

type TaskHandler struct {
	// `logger *zap.Logger` is a field of the `TaskHandler` struct that holds a reference to a logger
	// instance from the Zap logging library. This logger can be used to log messages and errors throughout
	// the codebase.
	logger                *zap.Logger
	postgresDb            *postgresdb.Db
	clickhouseDb          *clickhousedb.Db
	instrumentationClient *instrumentation.Client
	plaidClient           *plaidhandler.PlaidWrapper
}

var _ taskhandler.ITaskHandler = (*TaskHandler)(nil)

type Option func(*TaskHandler)

// The function WithLogger takes a pointer to a zap.Logger and returns an Option.
func WithLogger(logger *zap.Logger) Option {
	return func(handler *TaskHandler) {
		handler.logger = logger
	}
}

// The function takes a pointer to a Postgres database and returns an option.
func WithPostgresDb(postgresDb *postgresdb.Db) Option {
	return func(handler *TaskHandler) {
		handler.postgresDb = postgresDb
	}
}

// The function takes a pointer to a clickhousedb.Db object as input and returns an Option.
func WithClickhouseDb(clickhouseDb *clickhousedb.Db) Option {
	return func(handler *TaskHandler) {
		handler.clickhouseDb = clickhouseDb
	}
}

// This function takes an instrumentation client as input and returns an option.
func WithInstrumentationClient(instrumentationClient *instrumentation.Client) Option {
	return func(handler *TaskHandler) {
		handler.instrumentationClient = instrumentationClient
	}
}

// This function takes a Plaid client as input and returns an option.
func WithPlaidClient(plaidClient *plaidhandler.PlaidWrapper) Option {
	return func(handler *TaskHandler) {
		handler.plaidClient = plaidClient
	}
}

// This function returns a new instance of a TaskHandler with optional configuration options.
func NewTaskHandler(opts ...Option) *TaskHandler {
	handler := &TaskHandler{}

	for _, opt := range opts {
		opt(handler)
	}

	return handler
}

// `func (th *TaskHandler) Validate() error` is a method of the `TaskHandler` struct that is used to
// validate the fields of the `TaskHandler` instance. It returns an error if any of the required fields
// are missing or invalid. The method checks if the `TaskHandler` instance is not nil, if the
// `postgresDb` field is not nil, if the `clickhouseDb` field is not nil, and if the
// `instrumentationClient` field is not nil. If any of these fields are missing or invalid, the method
// returns an error indicating the specific issue.
func (th *TaskHandler) Validate() error {
	if th == nil {
		return ErrInvalidTaskHandler
	}

	if th.postgresDb == nil {
		return ErrInvalidPostgresConnectionObject
	}

	if th.clickhouseDb == nil {
		return ErrInvalidClickhouseConnectionObject
	}

	if th.instrumentationClient == nil {
		return ErrInvalidInstrumentationClient
	}

	if th.plaidClient == nil {
		return ErrInvalidPlaidClient
	}

	return nil
}

// `func (th *TaskHandler) RegisterTasks() *asynq.ServeMux` is a method of the `TaskHandler` struct
// that returns a new instance of `asynq.ServeMux`. This method is used to register task handlers for
// different types of tasks that can be executed asynchronously using the Asynq library. The
// `RegisterTasks` method also adds a middleware function `th.instrumentationMiddleware` to the
// `asynq.ServeMux` instance, which is used to emit traces and metrics to track the performance of the
// task handlers. In this specific implementation, the `RegisterTasks` method registers a task handler
// for the `process_transaction` task, but the actual implementation of the task handler function is
// not provided in the code snippet.
func (th *TaskHandler) RegisterTaskHandler() *asynq.ServeMux {
	mux := asynq.NewServeMux()
	mux.Use(th.instrumentationMiddleware)

	mux.HandleFunc(TaskSyncTransactions.String(), th.RunSyncPlaidTransactionsTask)
	mux.HandleFunc(TaskDeleteLink.String(), th.RunDeleteLinkTask)
	mux.HandleFunc(TaskDeleteTransactions.String(), th.RunDeleteTransactionsTask)
	mux.HandleFunc(TaskPullTransactions.String(), th.RunPullTransactionsTask)
	mux.HandleFunc(TaskPullUpdatedRecurringTransactions.String(), th.RunPullUpdatedReCurringTransactionsTask)
	mux.HandleFunc(TaskPullInvestmentTransactions.String(), th.RunPullInvestmentTransactionsTask)
	mux.HandleFunc(TaskPullInvestmentHoldings.String(), th.RunPullInvestmentHoldingsTask)
	mux.HandleFunc(TaskSyncNewLiabilityAccounts.String(), th.RunSyncNewLiabilityAccountsTask)
	return mux
}

// `func (th *TaskHandler) instrumentationMiddleware(next asynq.Handler) asynq.Handler` is a middleware
// function that takes in an `asynq.Handler` as an argument and returns an `asynq.Handler`. This
// middleware function is used to wrap around the actual task handler function and perform some
// additional actions before and after the task is executed. In this case, the middleware function is
// used to emit a trace to track the latency of the operation, increment and decrement task in-progress
// counter metric, increment task failed counter metric if the task fails, and increment task processed
// counter metric.
func (th *TaskHandler) instrumentationMiddleware(next asynq.Handler) asynq.Handler {
	return asynq.HandlerFunc(func(ctx context.Context, t *asynq.Task) error {
		// TODO: emit a trace here to track the latency of this given operation

		// TODO: increment task in-progress counter metric
		err := next.ProcessTask(ctx, t)
		// TODO: decrement task in progress counter metric
		if err != nil {
			// TODO: increment task failed counter metric
		}
		// TODO: increment task processed counter metric
		return err
	})
}
