package clickhousedatabase

import (
	"fmt"

	"github.com/SimifiniiCTO/simfiny-core-lib/database/clickhouse"
	"github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"
	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/api/v1"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/dal"
	"go.uber.org/zap"
)

const SAVE_POINT = "test_save_point"

var conn *Db

func init() {
	conn = NewTestDatabase()
}

func NewTestDatabase() *Db {
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
