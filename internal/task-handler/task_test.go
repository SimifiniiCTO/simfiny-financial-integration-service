package taskhandler

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/SimifiniiCTO/simfiny-core-lib/database/clickhouse"
	postgresdb "github.com/SimifiniiCTO/simfiny-core-lib/database/postgres"
	"github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"
	clickhousedatabase "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/clickhouse-database"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/plaidhandler"
	database "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/postgres-database"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/dal"
	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"go.uber.org/zap"
)

var (
	testTaskHandler *TaskHandler
	testAccessToken string
	testItemId      string
)

func setup() {
	clickhouseDb := NewClickhouseTestDatabase()
	postgresDb := NewPostgresTestDatabase()
	plaidTestClient, err := plaidhandler.GetPlaidWrapperForTest()
	if err != nil {
		log.Fatal(err)
	}

	testTaskHandler = &TaskHandler{
		logger:                zap.L(),
		postgresDb:            postgresDb,
		clickhouseDb:          clickhouseDb,
		instrumentationClient: &instrumentation.Client{},
		plaidClient:           plaidTestClient,
	}

	res, err := plaidTestClient.GetAccessTokenForSandboxAcct()
	if err != nil {
		log.Fatal(err)
	}

	testAccessToken = res.AccessToken
	testItemId = res.ItemId
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func NewClickhouseTestDatabase() *clickhousedatabase.Db {
	client, err := clickhouse.NewInMemoryTestDbClient(schema.GetClickhouseSchemas()...)
	if err != nil {
		panic(fmt.Errorf("failed to create in memory test db client: %w", err))
	}

	return &clickhousedatabase.Db{
		Conn:                  client,
		QueryOperator:         dal.Use(client.Engine),
		Logger:                zap.NewNop(),
		InstrumentationClient: &instrumentation.Client{},
	}
}

func NewPostgresTestDatabase() *database.Db {
	client, err := postgresdb.NewInMemoryTestDbClient(schema.GetDatabaseSchemas()...)
	if err != nil {
		panic(fmt.Errorf("failed to create in memory test db client: %w", err))
	}

	return &database.Db{
		Conn:            client,
		QueryOperator:   dal.Use(client.Engine),
		Logger:          zap.NewNop(),
		Instrumentation: &instrumentation.Client{},
	}
}
