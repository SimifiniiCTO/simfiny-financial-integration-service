package transactionmanager

import (
	"context"
	"fmt"
	"time"

	"go.temporal.io/api/serviceerror"
	"go.temporal.io/api/workflowservice/v1"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/database"
)

var (
	ErrInvalidClientOptions                = fmt.Errorf("invalid input argument. client options cannot be null")
	ErrInvalidTransactionManagerConfig     = fmt.Errorf("invalid input argument. transaction manager config cannot be nil")
	ErrInvalidLogger                       = fmt.Errorf("invalid logger. logger cannot be nil")
	ErrInvalidTelemetrySDK                 = fmt.Errorf("invalid telemetry sdk config. config cannot be nil")
	ErrNilAccount                          = fmt.Errorf("account cannot be nil")
	ErrNilOldEmail                         = fmt.Errorf("old email cannot be nil")
	ErrInvalidUpdateAccountWorkflowRequest = fmt.Errorf("invalid update account workflow request. request cannot be nil or invalid")
	ErrInvalidRetryPolicy                  = fmt.Errorf("invalid input argument. retry policy cannot be nil or invalid")
	ErrInvalidRpcTimeout                   = fmt.Errorf("invalid input argument. rpc timeout cannot be nil or invalid")
)

// `TransactionManager` is the single struct by which we manage and initiate all distributed transactions
// within the service. It provides wrapper facilities around the temporal sdk client as well in order
// to properly emit metrics and traces during rpc operations
//
// @property client - This is the client used to interact with a remote temporal cluster.
// @property TelemetrySDK - This is the telemetry SDK that we use to send telemetry data to newrelic
// @property Logger - This is the logger that will be used to log messages.
// Service.
// @property DatabaseConn - This is the database connection object that we will use to connect to the
// database.
type TransactionManager struct {
	Client                             client.Client
	NamespaceClient                    client.NamespaceClient
	TelemetrySDK                       *instrumentation.Client
	Logger                             *zap.Logger
	DatabaseConn                       *database.Db
	RetryPolicy                        *Policy
	RPCTimeout                         time.Duration
	Worker                             worker.Worker
	MetricsEnabled                     bool
	MaxConcurrentActivityExecutionSize uint64
	MaxConcurrentWorkflowTaskPollers   uint64
	ServiceTaskQueue                   string
}

// Policy outlines retry policies necessary for workflow and downstream service calls
type Policy struct {
	RetryInitialInterval    *time.Duration
	RetryBackoffCoefficient float64
	MaximumInterval         time.Duration
	MaximumAttempts         int
}

type WorkflowManager interface {
	// DeleteBankAccountWorkflow use do delete a bank account initially across all participant services
	// @param ctx - This is the workflow context.
	// @param bankAccountID - This is the bank account id of the bank account to be deleted.
	// @return - This returns an error if the operation fails.
	DeleteBankAccountWorkflow(ctx workflow.Context, bankAccountID uint64) error

	// DeleteProfileWorkflow use do delete a profile initially across all participant services
	// @param ctx - This is the workflow context.
	// @param profileID - This is the profile id of the profile to be deleted.
	// @return - This returns an error if the operation fails.
	DeleteProfileWorkflow(ctx workflow.Context, profileID uint64) error
}

type WorkflowActivityManager interface {
	// DeleteBankAccountActivity use do delete a bank account initially across all participant services
	// @param ctx - This is the activity context.
	// @param bankAccountID - This is the bank account id of the bank account to be deleted.
	// @return - This returns an error if the operation fails.
	DeleteBankAccountActivity(ctx context.Context, bankAccountID uint64) error

	// Delete Profile Activity use do delete a profile initially across all participant services
	// @param ctx - This is the activity context.
	// @param profileID - This is the profile id of the profile to be deleted.
	// @return - This returns an error if the operation fails.
	DeleteProfileActivity(ctx context.Context, profileID uint64) error
}

var _ WorkflowManager = &TransactionManager{}
var _ WorkflowActivityManager = &TransactionManager{}

// It creates a new instance of the TransactionManager struct and returns it
func NewTransactionManager(opts ...Option) (*TransactionManager, error) {
	configs := &Configs{}

	for _, opt := range opts {
		opt(configs)
	}

	if err := configs.validate(); err != nil {
		return nil, err
	}

	l := configs.Logger
	l.Info("creating new transaction manager", zap.Any("configs", configs.ClientOptions.HostPort))

	temporalClient, err := newClient(configs.ClientOptions)
	if err != nil {
		l.Error("failed to create temporal client", zap.Error(err))
		return nil, err
	}

	nsClient, err := newNamespaceClient(configs.ClientOptions)
	if err != nil {
		l.Error("failed to create temporal namespace client", zap.Error(err))
		return nil, err
	}

	txManager := &TransactionManager{
		Client:                             temporalClient,
		NamespaceClient:                    nsClient,
		TelemetrySDK:                       configs.TelemetrySDK,
		Logger:                             configs.Logger,
		DatabaseConn:                       configs.DatabaseConn,
		RetryPolicy:                        configs.RetryPolicy,
		RPCTimeout:                         *configs.RpcTimeout,
		MetricsEnabled:                     configs.MetricsEnabled,
		MaxConcurrentActivityExecutionSize: configs.MaxConcurrentActivityExecutionSize,
		MaxConcurrentWorkflowTaskPollers:   configs.MaxConcurrentWorkflowTaskPollers,
		ServiceTaskQueue:                   configs.ServiceTaskQueue,
	}

	l.Info("transaction manager created successfully")

	txManager.initializeWorker()
	return txManager, nil
}

func (tx *TransactionManager) initializeWorker() {
	l := tx.Logger
	tx.Worker = worker.New(tx.Client, tx.ServiceTaskQueue, worker.Options{
		// TODO: research below fields and ensure we properly populate them
		MaxConcurrentActivityExecutionSize: int(tx.MaxConcurrentActivityExecutionSize),
		MaxConcurrentWorkflowTaskPollers:   int(tx.MaxConcurrentWorkflowTaskPollers),
	})

	l.Info("initializing worker", zap.String("task_queue", tx.ServiceTaskQueue))
	l.Info("registering workflows and activities")

	// register workflow
	tx.Worker.RegisterWorkflow(tx.DeleteBankAccountWorkflow)
	tx.Worker.RegisterWorkflow(tx.DeleteProfileWorkflow)

	// register activity
	tx.Worker.RegisterActivity(tx.DeleteBankAccountActivity)
	tx.Worker.RegisterActivity(tx.DeleteProfileActivity)

	l.Info("worker initialized")
}

// StartWorker enables the worker to start listening to a given task queue
// This should ideally be ran in a go routine
func (tx *TransactionManager) StartWorker() {
	err := tx.Worker.Run(worker.InterruptCh())
	if err != nil {
		tx.Logger.Fatal("unable to start Worker", zap.Error(err))
	}
}

func (t *TransactionManager) Close() {
	t.Client.Close()
	t.NamespaceClient.Close()
}

// Instantiates a new client
func newClient(opts *client.Options) (client.Client, error) {
	if opts == nil {
		return nil, ErrInvalidClientOptions
	}
	// Create the client object just once per process
	c, err := client.Dial(*opts)
	if err != nil {
		return nil, err
	}

	return c, nil
}

// NewNamespaceClient creates an instance of a namespace client, to manage
// lifecycle of namespaces.
func newNamespaceClient(opts *client.Options) (client.NamespaceClient, error) {
	if opts == nil {
		return nil, ErrInvalidClientOptions
	}

	// calls the initialize a new namespace will not attempt to connect to the
	// temporal cluster eagerly hence the call may not fail even if the server is unreachable
	// we need to pass grpc.WithBlock as a gRPC dial option to connection options to eagerly connect
	connectionOptions := opts.ConnectionOptions.DialOptions
	connectionOptions = append(connectionOptions, grpc.WithBlock())
	opts.ConnectionOptions.DialOptions = connectionOptions

	// Create the client object just once per process
	c, err := client.NewNamespaceClient(*opts)
	if err != nil {
		return nil, err
	}

	// upon successfully creating the namespace client, we ensure to also create the namespace
	// on which our workers will process tasks
	// TODO: read this from env variables
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	workflowRetentionPeriod := time.Hour * 24
	if err := c.Register(ctx, &workflowservice.RegisterNamespaceRequest{
		Namespace:                        opts.Namespace,
		Description:                      "simfiny financial integration service namespace",
		OwnerEmail:                       "yoan@simfinii.com",
		WorkflowExecutionRetentionPeriod: &workflowRetentionPeriod,
	}); err != nil {
		// ignore the error if the namespace already exists
		if _, ok := err.(*serviceerror.NamespaceAlreadyExists); !ok {
			return nil, err
		}
	}

	return c, nil
}
