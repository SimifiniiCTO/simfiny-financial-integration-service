package clickhousedatabase

import (
	"context"
	"os"
	"time"

	"github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"
	"github.com/uptrace/go-clickhouse/ch"
	"github.com/uptrace/go-clickhouse/chdebug"
	"go.uber.org/zap"
)

const SAVE_POINT = "test_save_point"

var conn *Db

func init() {
	conn = NewTestDatabase()
}

func NewTestDatabase() *Db {
	// client, err := clickhouse.NewInMemoryTestDbClient(schema.GetClickhouseSchemas()...)
	// if err != nil {
	// 	panic(fmt.Errorf("failed to create in memory test db client: %w", err))
	// }

	opts := []ch.Option{
		ch.WithTimeout(5 * time.Second),
		ch.WithDialTimeout(5 * time.Second),
		ch.WithReadTimeout(5 * time.Second),
		ch.WithWriteTimeout(5 * time.Second),
	}

	db := chDB(opts...)

	orm := &Db{
		Conn:                  nil,
		QueryOperator:         nil, // dal.Use(client.Engine),
		Logger:                zap.NewNop(),
		InstrumentationClient: &instrumentation.Client{},
		queryEngine:           db,
	}

	if err := orm.performSchemaMigration(context.Background()); err != nil {
		panic(err)
	}

	return orm
}

func chDB(opts ...ch.Option) *ch.DB {
	dsn := os.Getenv("CH")
	if dsn == "" {
		dsn = "clickhouse://gorm:gorm@localhost:9942/gorm?dial_timeout=10s&read_timeout=20s&sslmode=disable"
	}

	opts = append(opts, ch.WithDSN(dsn), ch.WithAutoCreateDatabase(true))
	db := ch.Connect(opts...)
	db.AddQueryHook(chdebug.NewQueryHook(
		chdebug.WithEnabled(true),
		chdebug.WithVerbose(true),
	))
	return db
}
