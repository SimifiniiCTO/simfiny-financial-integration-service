package clickhousedatabase

import (
	"context"
	"fmt"
	"time"

	"github.com/SimifiniiCTO/simfiny-core-lib/database/clickhouse"
	clickhousedb "github.com/SimifiniiCTO/simfiny-core-lib/database/clickhouse"
	"github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/migrations"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/service_errors"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/dal"
	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/uptrace/go-clickhouse/ch"
	chgo "github.com/uptrace/go-clickhouse/ch"
	"go.uber.org/zap"
)

type ClickhouseDatabaseOperations interface {
	// `AddTransaction` is a method defined in the `ClickhouseDatabaseOperations` interface in the Go
	// programming language. It takes in a context and an interface representing a transaction, and returns
	// an error. This method is used to add a single transaction to a Clickhouse database.
	AddTransaction(ctx context.Context, userId *uint64, tx *schema.Transaction) (*string, error)

	// `AddTransactions` is a method defined in the `ClickhouseDatabaseOperations` interface in the Go
	// programming language. It takes in a context and a slice of interfaces representing transactions, and
	// returns an error. This method is used to add multiple transactions to a Clickhouse database.
	AddTransactions(ctx context.Context, userId *uint64, txs []*schema.Transaction) error

	// // The `UpdateTransaction` method is defined in the `ClickhouseDatabaseOperations` interface in the Go
	// // programming language. It takes in a context, a user ID as a uint64, a transaction ID as a uint64,
	// // and an interface representing a transaction, and returns an error. This method is used to update a
	// // single transaction with the specified transaction ID for a specific user in a Clickhouse database.
	// UpdateTransaction(ctx context.Context, userId *uint64, txId *uint64, tx *schema.Transaction) error

	// The `DeleteTransaction` method is defined in the `ClickhouseDatabaseOperations` interface in the Go
	// programming language. It takes in a context and a transaction ID as a uint64, and returns an error.
	// This method is used to delete a single transaction with the specified transaction ID from a
	// Clickhouse database.
	DeleteTransaction(ctx context.Context, txId *string) error

	// The `DeleteTransactons` method is defined in the `ClickhouseDatabaseOperations` interface in the Go
	// programming language. It takes in a context and a user ID as a uint64, and returns an error. This
	// method is used to delete all transactions associated with a specific user from a Clickhouse
	// database.
	DeleteUserTransactions(ctx context.Context, userId *uint64) error

	// The `DeleteTransactionsByIds` method is defined in the `ClickhouseDatabaseOperations` interface in
	// the Go programming language. It takes in a context and a variable number of transaction IDs as
	// uint64 values, and returns an error. This method is used to delete multiple transactions with the
	// specified transaction IDs from a Clickhouse database.
	// DeleteTransactionsByIds(ctx context.Context, txIds []string) error

	// `GetTransactions` is a method defined in the `ClickhouseDatabaseOperations` interface in the Go
	// programming language. It takes in a context and a user ID as a uint64, and returns a slice of
	// interfaces representing transactions and an error. This method is used to retrieve all transactions
	// associated with a specific user from a Clickhouse database.
	GetTransactions(context.Context, *uint64, int64, int64) ([]*schema.Transaction, int64, error)
	// `GetTransactionById` is a method defined in the `ClickhouseDatabaseOperations` interface in the Go
	// programming language. It takes in a context and a transaction ID as a uint64, and returns a pointer
	// to a `schema.Transaction` struct and an error. This method is used to retrieve a single transaction
	// with the specified transaction ID from a Clickhouse database.
	// GetTransactionById(ctx context.Context, txId *uint64) (*schema.Transaction, error)
}

type Db struct {
	connectionUri string
	queryEngine   *chgo.DB
	// `Conn *clickhousedb.Client` is defining a field named `Conn` of type `*clickhousedb.Client`. This
	// field is used to store a connection to a Clickhouse database.
	Conn *clickhousedb.Client
	// `logger *zap.Logger` is defining a field named `logger` of type `*zap.Logger`. This field is used to
	// store an instance of the `zap.Logger` struct, which is a popular logging library in the Go
	// programming language. It is used to log messages and events during the execution of the code.
	Logger *zap.Logger

	// `instrumentationClient *instrumentation.Client` is defining a field named `instrumentationClient` of
	// type `*instrumentation.Client`. This field is used to store an instance of the
	// `instrumentation.Client` struct, which is used for collecting and reporting metrics and tracing data
	// during the execution of the code. It is likely used to monitor the performance and behavior of the
	// Clickhouse database operations.
	InstrumentationClient *instrumentation.Client

	// `queryOperator *dal.Query` is defining a field named `queryOperator` of type `*dal.Query`. This
	// field is used to store an instance of the `dal.Query` struct, which is a generated data access layer
	// (DAL) that provides methods for interacting with a Clickhouse database. It is likely used to execute
	// queries and retrieve data from the database.
	QueryOperator *dal.Query
}

var _ ClickhouseDatabaseOperations = (*Db)(nil)

// New returns a new database object
func New(ctx context.Context, opts ...Option) (*Db, error) {
	database := &Db{}

	for _, opt := range opts {
		opt(database)
	}

	// validate the database object
	if err := database.Validate(); err != nil {
		return nil, err
	}

	database.Logger.Info("successfully validated clickhouse db")
	database.Logger.Info(database.connectionUri)
	db := ch.Connect(
		// clickhouse://<user>:<password>@<host>:<port>/<database>?sslmode=disable
		ch.WithDSN(database.connectionUri),
		ch.WithTimeout(5*time.Second),
		ch.WithDialTimeout(5*time.Second),
		ch.WithReadTimeout(5*time.Second),
		ch.WithWriteTimeout(5*time.Second),
		// ch.WithTLSConfig(&tls.Config{InsecureSkipVerify: true}),
	)

	if db != nil {
		database.Logger.Info("successful connection to database with uptrace client")
	}

	// ping the database
	if err := db.Ping(ctx); err != nil {
		return nil, err
	}

	database.Logger.Info("successful pinged database with uptrace")
	database.queryEngine = db

	// perform migrations
	if err := database.performSchemaMigration(ctx); err != nil {
		return nil, err
	}

	return database, nil
}

// The `performSchemaMigration` function is a method defined on the `Db` struct. It takes a context as
// an argument and returns an error. This method is responsible for performing schema migration on the
// Clickhouse database.
func (db *Db) performSchemaMigration(ctx context.Context) error {
	if db == nil {
		return service_errors.ErrInvalidAcctParam
	}

	db.Logger.Info("Defining migration engine")
	migationEngine, err := migrations.New(db.queryEngine)
	if err != nil {
		db.Logger.Error(err.Error())
		return err
	}

	// TODO: NEED TO PROPERLY SUPPORT MIGRATIONS (REF: https://github.com/go-gormigrate/gormigrate)
	// 	// ref. https://kb.altinity.com/engines/mergetree-table-engine-family/collapsing-vs-replacing/
	// 	// ref. https://clickhouse.com/docs/en/guides/developer/deduplication
	db.Logger.Info("Performing database migrations")
	if err := migationEngine.Migrate(ctx); err != nil {
		db.Logger.Error(err.Error())
		return err
	}

	db.Logger.Info("Successfully migrated clickhouse database schemas")
	return nil
}

// startDatastoreSpan generates a datastore span
func (db *Db) startDatastoreSpan(ctx context.Context, name string) *newrelic.DatastoreSegment {
	if db.InstrumentationClient != nil {
		txn := db.InstrumentationClient.GetTraceFromContext(ctx)
		span := db.InstrumentationClient.StartDatastoreSegment(txn, name)
		return span
	}

	return nil
}

// NewMockInMemoryClickhouseDB returns a new mock in memory clickhouse database
func NewMockInMemoryClickhouseDB() *Db {
	client, err := clickhouse.NewInMemoryTestDbClient(schema.GetClickhouseSchemas()...)
	if err != nil {
		panic(fmt.Errorf("failed to create in memory test db client: %w", err))
	}

	return &Db{
		Conn:                  client,
		QueryOperator:         dal.Use(client.Engine),
		Logger:                zap.NewNop(),
		InstrumentationClient: &instrumentation.Client{},
	}
}
