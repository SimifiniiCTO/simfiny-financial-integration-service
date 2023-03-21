package database

import (
	"context"
	"fmt"
	"sync"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	core_database "github.com/SimifiniiCTO/core/core-database"
	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/dal"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/secrets"
)

const dbName = "gen_test.db"

var testdb *gorm.DB
var once sync.Once
var conn *Db

type ShimKMS struct {
}

var _ secrets.KeyManagement = &ShimKMS{}

func (s *ShimKMS) Decrypt(ctx context.Context, keyID, version string, input []byte) (result []byte, _ error) {
	data := []byte("sldjfdshfljd")
	return data, nil
}

func (s *ShimKMS) Encrypt(ctx context.Context, input []byte) (keyID, version string, result []byte, _ error) {
	result = []byte("sldjfdshfljd")
	keyID = "kshdkfjsdfsdfd"
	version = "1"
	return keyID, version, result, nil
}

func init() {
	conn = NewTestDatabase()
}

func NewTestDatabase() *Db {
	InitializeDB()
	testdb.AutoMigrate(schema.GetDatabaseSchemas()...)
	return &Db{
		Conn:                   &core_database.DatabaseConn{Engine: testdb},
		Logger:                 zap.L(),
		MaxConnectionAttempts:  3,
		MaxRetriesPerOperation: 0,
		RetryTimeOut:           1 * time.Minute,
		OperationSleepInterval: 500 * time.Millisecond,
		QueryOperator:          dal.Use(testdb),
		Kms:                    nil,
	}
}

func InitializeDB() {
	once.Do(func() {
		var err error
		testdb, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
		if err != nil {
			panic(fmt.Errorf("open sqlite %q fail: %w", dbName, err))
		}
	})
}
