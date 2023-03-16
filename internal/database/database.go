package database

import (
	"context"
	"time"

	core_database "github.com/yoanyombapro1234/FeelGuuds_Core/core/core-database"
	"go.uber.org/zap"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/generated/api/v1"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/generated/dal"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/instrumentation"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/service_errors"
)

// DatabaseOperations provides an interface which any database tied to this service should implement
type DatabaseOperations interface {
	// Account Operations
	// CreateUserProfile creates a user profile
	CreateUserProfile(ctx context.Context, profile *schema.UserProfile) (*schema.UserProfile, error)
	// DeleteUserProfileByUserId deletes a user profile by user id
	DeleteUserProfileByUserID(ctx context.Context, userID uint64) error
	// UpdateUserProfile updates a user profile
	UpdateUserProfile(ctx context.Context, profile *schema.UserProfile) error
	// GetUserProfileByUserID retrieves a user profile by user id
	GetUserProfileByUserID(ctx context.Context, userID uint64) (*schema.UserProfile, error)
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
	Instrumentation *instrumentation.ServiceTelemetry
	// QueryOperator is the object that will be used to execute database queries
	QueryOperator *dal.Query
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
	Instrumentation *instrumentation.ServiceTelemetry
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
	databaseModels := schema.GetDatabaseSchemas()

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
		Instrumentation:        params.Instrumentation,
		QueryOperator:          dal.Use(conn.Engine),
	}, nil
}
