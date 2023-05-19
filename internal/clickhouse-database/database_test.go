package clickhousedatabase

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/SimifiniiCTO/simfiny-core-lib/database/clickhouse"
	"github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"
	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/dal"
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
		queryOperator:         dal.Use(client.Engine),
		logger:                zap.NewNop(),
		instrumentationClient: &instrumentation.Client{},
	}
}
