package clickhousedatabase

import (
	"context"
	"fmt"

	"github.com/SimifiniiCTO/simfiny-core-lib/database/clickhouse"
	clickhousedb "github.com/SimifiniiCTO/simfiny-core-lib/database/clickhouse"
	"github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/service_errors"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/dal"
	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"github.com/newrelic/go-agent/v3/newrelic"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ClickhouseDatabaseOperations interface {
	// `AddTransaction` is a method defined in the `ClickhouseDatabaseOperations` interface in the Go
	// programming language. It takes in a context and an interface representing a transaction, and returns
	// an error. This method is used to add a single transaction to a Clickhouse database.
	AddTransaction(ctx context.Context, userId *uint64, tx *schema.Transaction) (*uint64, error)

	// `AddTransactions` is a method defined in the `ClickhouseDatabaseOperations` interface in the Go
	// programming language. It takes in a context and a slice of interfaces representing transactions, and
	// returns an error. This method is used to add multiple transactions to a Clickhouse database.
	AddTransactions(ctx context.Context, userId *uint64, txs []*schema.Transaction) error

	// The `UpdateTransaction` method is defined in the `ClickhouseDatabaseOperations` interface in the Go
	// programming language. It takes in a context, a user ID as a uint64, a transaction ID as a uint64,
	// and an interface representing a transaction, and returns an error. This method is used to update a
	// single transaction with the specified transaction ID for a specific user in a Clickhouse database.
	UpdateTransaction(ctx context.Context, userId *uint64, txId *uint64, tx *schema.Transaction) error

	// The `DeleteTransaction` method is defined in the `ClickhouseDatabaseOperations` interface in the Go
	// programming language. It takes in a context and a transaction ID as a uint64, and returns an error.
	// This method is used to delete a single transaction with the specified transaction ID from a
	// Clickhouse database.
	DeleteTransaction(ctx context.Context, txId *uint64) error

	// The `DeleteTransactons` method is defined in the `ClickhouseDatabaseOperations` interface in the Go
	// programming language. It takes in a context and a user ID as a uint64, and returns an error. This
	// method is used to delete all transactions associated with a specific user from a Clickhouse
	// database.
	DeleteUserTransactons(ctx context.Context, userId *uint64) error

	// The `DeleteTransactionsByIds` method is defined in the `ClickhouseDatabaseOperations` interface in
	// the Go programming language. It takes in a context and a variable number of transaction IDs as
	// uint64 values, and returns an error. This method is used to delete multiple transactions with the
	// specified transaction IDs from a Clickhouse database.
	DeleteTransactionsByIds(ctx context.Context, txIds []uint64) error

	// `GetTransactions` is a method defined in the `ClickhouseDatabaseOperations` interface in the Go
	// programming language. It takes in a context and a user ID as a uint64, and returns a slice of
	// interfaces representing transactions and an error. This method is used to retrieve all transactions
	// associated with a specific user from a Clickhouse database.
	GetTransactions(context.Context, *uint64, int64, int64) ([]*schema.Transaction, int64, error)
	// `GetTransactionById` is a method defined in the `ClickhouseDatabaseOperations` interface in the Go
	// programming language. It takes in a context and a transaction ID as a uint64, and returns a pointer
	// to a `schema.Transaction` struct and an error. This method is used to retrieve a single transaction
	// with the specified transaction ID from a Clickhouse database.
	GetTransactionById(ctx context.Context, txId *uint64) (*schema.Transaction, error)
}

type Db struct {
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

	// ping the database
	if err := database.pingDatabase(); err != nil {
		return nil, err
	}

	// perform migrations
	if err := database.performSchemaMigration(); err != nil {
		return nil, err
	}

	return database, nil
}

func (db *Db) pingDatabase() error {
	if db == nil {
		return service_errors.ErrInvalidAcctParam
	}

	conn, err := db.Conn.Engine.DB()
	if err != nil {
		return err
	}

	if err := conn.Ping(); err != nil {
		return err
	}

	return nil
}

func (db *Db) performSchemaMigration() error {
	var (
		engine *gorm.DB
		models = schema.GetClickhouseSchemas()
	)

	if db == nil {
		return service_errors.ErrInvalidAcctParam
	}

	if engine = db.Conn.Engine; engine == nil {
		return service_errors.ErrInvalidGormDbOject
	}

	// TODO: NEED TO PROPERLY SUPPORT MIGRATIONS (REF: https://github.com/go-gormigrate/gormigrate)
	if len(models) > 0 {
		// ref. https://kb.altinity.com/engines/mergetree-table-engine-family/collapsing-vs-replacing/
		// ref. https://clickhouse.com/docs/en/guides/developer/deduplication
		for _, e := range models {
			if !engine.Migrator().HasTable(e) {
				if err := engine.AutoMigrate(e); err != nil {
					// TODO: emit failure metric here
					// money
					db.Logger.Error(err.Error())
				}
			}

		}
		db.Logger.Info("successfully migrated database schemas")
	}

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
