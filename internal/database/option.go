package database

import (
	"time"

	"go.uber.org/zap"

	core_database "github.com/SimifiniiCTO/core/core-database"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/dal"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/instrumentation"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/secrets"
)

// Option is a function that configures a Db
type Option func(*Db)

// WithDatabaseLogger configures the database logger
func WithDatabaseLogger(logger *zap.Logger) Option {
	return func(d *Db) {
		d.Logger = logger
	}
}

// WithDatabaseMaxConnectionAttempts configures the database max connection attempts
func WithDatabaseMaxConnectionAttempts(maxConnectionAttempts int) Option {
	return func(d *Db) {
		d.MaxConnectionAttempts = maxConnectionAttempts
	}
}

// WithDatabaseMaxRetriesPerOperation configures the database max retries per operation
func WithDatabaseMaxRetriesPerOperation(maxRetriesPerOperation int) Option {
	return func(d *Db) {
		d.MaxRetriesPerOperation = maxRetriesPerOperation
	}
}

// WithDatabaseRetryTimeOut configures the database retry timeout
func WithDatabaseRetryTimeOut(retryTimeOut time.Duration) Option {
	return func(d *Db) {
		d.RetryTimeOut = retryTimeOut
	}
}

// WithDatabaseOperationSleepInterval configures the database operation sleep interval
func WithDatabaseOperationSleepInterval(operationSleepInterval time.Duration) Option {
	return func(d *Db) {
		d.OperationSleepInterval = operationSleepInterval
	}
}

// WithDatabaseInstrumentation configures the database instrumentation
func WithDatabaseInstrumentation(instrumentation *instrumentation.ServiceTelemetry) Option {
	return func(d *Db) {
		d.Instrumentation = instrumentation
	}
}

// WithDatabaseQueryOperator configures the database query operator
func WithDatabaseQueryOperator(queryOperator *dal.Query) Option {
	return func(d *Db) {
		d.QueryOperator = queryOperator
	}
}

// WithConnection configures the database connection
func WithConnection(params *core_database.DatabaseConn) Option {
	return func(d *Db) {
		d.Conn = params
	}
}

// WithKms configures the database kms
func WithKms(kms secrets.KeyManagement) Option {
	return func(d *Db) {
		d.Kms = kms
	}
}
