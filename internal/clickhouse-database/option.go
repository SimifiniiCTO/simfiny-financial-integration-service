package clickhousedatabase

import (
	clickhousedb "github.com/SimifiniiCTO/simfiny-core-lib/database/clickhouse"
	"github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/dal"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/service_errors"
	"go.uber.org/zap"
)

type Option func(*Db)

// WithDatabaseInstrumentation configures the database instrumentation
func WithDatabaseInstrumentation(instrumentation *instrumentation.Client) Option {
	return func(d *Db) {
		d.InstrumentationClient = instrumentation
	}
}

// WithDatabaseQueryOperator configures the database query operator
func WithDatabaseQueryOperator(queryOperator *dal.Query) Option {
	return func(d *Db) {
		d.QueryOperator = queryOperator
	}
}

// WithDatabaseLogger configures the database logger
func WithDatabaseLogger(logger *zap.Logger) Option {
	return func(d *Db) {
		d.Logger = logger
	}
}

// WithDatabaseClient returns a new database client
func WithDatabaseClient(client *clickhousedb.Client) Option {
	return func(d *Db) {
		d.Conn = client
	}
}

// Validate validates the database object
func (db *Db) Validate() error {
	if db.Conn == nil {
		return service_errors.ErrInvalidDbObject
	}

	if db.Logger == nil {
		return service_errors.ErrInvalidDbObject
	}

	if db.InstrumentationClient == nil {
		return service_errors.ErrInvalidDbObject
	}

	if db.QueryOperator == nil {
		return service_errors.ErrInvalidDbObject
	}

	if db.Conn == nil {
		return service_errors.ErrInvalidConn
	}

	return nil
}
