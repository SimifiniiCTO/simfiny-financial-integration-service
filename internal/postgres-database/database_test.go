package database

import (
	"context"
	"fmt"
	"sync"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"

	postgresdb "github.com/SimifiniiCTO/simfiny-core-lib/database/postgres"
	"github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"
	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/dal"
)

const dbName = "gen_test.db"
const SAVE_POINT = "test_save_point"

var testdb *gorm.DB
var once sync.Once
var conn *Db

func init() {
	conn = NewTestDatabase()
}

func NewTestDatabase() *Db {
	client, err := postgresdb.NewInMemoryTestDbClient(schema.GetDatabaseSchemas()...)
	if err != nil {
		panic(fmt.Errorf("failed to create in memory test db client: %w", err))
	}

	return &Db{
		Conn:            client,
		QueryOperator:   dal.Use(client.Engine),
		Logger:          zap.NewNop(),
		Instrumentation: &instrumentation.Client{},
	}
}

type TxCleanupHandler struct {
	cancelFunc               context.CancelFunc
	Tx                       *gorm.DB
	savePointRollbackHandler func(tx *gorm.DB)
}

func txCleanupHandler() *TxCleanupHandler {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	db := conn.Conn.Engine
	tx := db.WithContext(ctx).Begin()
	tx.SavePoint(SAVE_POINT)

	return &TxCleanupHandler{
		cancelFunc: cancel,
		Tx:         tx,
		savePointRollbackHandler: func(tx *gorm.DB) {
			tx.RollbackTo(SAVE_POINT)
			tx.Commit()
		},
	}
}
