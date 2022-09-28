package database

import (
	"context"
	"time"

	"github.com/SimifiniiCTO/simfinii/src/backend/services/financial-integration-service/internal/service_errors"
	schema "github.com/SimifiniiCTO/simfinii/src/backend/services/financial-integration-service/proto"
	"github.com/newrelic/go-agent/v3/newrelic"
	core_database "github.com/yoanyombapro1234/FeelGuuds_Core/core/core-database"
	"go.uber.org/zap"
)

// DatabaseOperations provides an interface which any database tied to this service should implement
type DatabaseOperations interface {
	FindVirtualAcctID(ctx context.Context, vAcctID uint64) (*schema.VirtualAccount, error)
	findVirtualAcctByIDTxn(ctx context.Context, txn *newrelic.Transaction, vAcctID uint64) core_database.CmplxTx
	FindVirtualAcctByUserID(ctx context.Context, userID uint64) (*schema.VirtualAccount, error)
	findVirtualAcctByUserIDTxn(ctx context.Context, txn *newrelic.Transaction, userID uint64) core_database.CmplxTx
	CreateVirtualAccount(ctx context.Context, vAcct *schema.VirtualAccount, accessToken string) (*schema.VirtualAccount, error)
	createVirtualAccountTxn(ctx context.Context, txn *newrelic.Transaction, vAcct *schema.VirtualAccount, accessToken string) core_database.CmplxTx
	DeactivateVirtualAccount(ctx context.Context, vAcctID uint64) error
	deactivateVirtualAccountTxn(ctx context.Context, txn *newrelic.Transaction, vAcctID uint64) core_database.Tx
	SaveVirtualAccountRecord(ctx context.Context, vAcct *schema.VirtualAccount) error
	saveVirtualAccountTxFunc(ctx context.Context, txn *newrelic.Transaction, vAcct *schema.VirtualAccount) core_database.Tx
	UpdateVirtualAccount(ctx context.Context, vAcct *schema.VirtualAccount, vAcctID uint64) error
	updateVirtualAccountTxn(ctx context.Context, txn *newrelic.Transaction, vAcct *schema.VirtualAccount, vAcctID uint64) core_database.Tx
}

// Db withholds connection to a postgres database as well as a logging handler
type Db struct {
	// Conn serves as the actual database connection object
	Conn *core_database.DatabaseConn
	// Logger is the logging utility used by this object
	Logger *zap.Logger
	// MaxConnectionAttempts outlines the maximum connection attempts
	// to initiate against the database
	MaxConnectionAttempts int
	// MaxRetriesPerOperation defines the maximum retries to attempt per failed database
	// connection attempt
	MaxRetriesPerOperation int
	// RetryTimeOut defines the maximum time until a retry operation is observed as a
	// timed out operation
	RetryTimeOut time.Duration
	// OperationSleepInterval defines the amount of time between retry operations
	// that the system sleeps
	OperationSleepInterval time.Duration
	// Telemetry defines the object by which we will emit metrics, trace requests, and database operations
	Telemetry *newrelic.Application
}

var _ DatabaseOperations = (*Db)(nil)

// ConnectionInitializationParams represents connection initialization parameters for the database
type ConnectionInitializationParams struct {
	// ConnectionParams outlines database connection parameters
	ConnectionParams *ConnectionParameters
	// Logger is the logging utility used by this object
	Logger *zap.Logger
	// MaxConnectionAttempts outlines the maximum connection attempts
	// to initiate against the database
	MaxConnectionAttempts int
	// MaxRetriesPerOperation defines the maximum retries to attempt per failed database
	// connection attempt
	MaxRetriesPerOperation int
	// RetryTimeOut defines the maximum time until a retry operation is observed as a
	// timed out operation
	RetryTimeOut time.Duration
	// RetrySleepInterval defines the amount of time between retry operations
	// that the system sleeps
	RetrySleepInterval time.Duration
	// Telemetry defines the object by which we will emit metrics, trace requests, and database operations
	Telemetry *newrelic.Application
}

// New creates a database connection and returns the connection object
func New(ctx context.Context, params *ConnectionInitializationParams) (*Db,
	error) {
	errInvalidInputParams := service_errors.ErrInvalidInputParam

	// TODO: generate a span for the database connection attempt
	if params == nil || (params != nil && (params.ConnectionParams == nil || params.Logger == nil)) {
		return nil, errInvalidInputParams
	}

	logger := params.Logger
	databaseModels := schema.Schemas()

	conn, err := connectToDatabase(ctx, params.ConnectionParams, params.Logger, databaseModels...)

	if err != nil {
		return nil, err
	}

	logger.Info("Successfully connected to the database")

	return &Db{
		Conn:                   conn,
		Logger:                 logger,
		MaxConnectionAttempts:  params.MaxConnectionAttempts,
		MaxRetriesPerOperation: params.MaxRetriesPerOperation,
		RetryTimeOut:           params.RetryTimeOut,
		OperationSleepInterval: params.RetrySleepInterval,
		Telemetry:              params.Telemetry,
	}, nil
}
