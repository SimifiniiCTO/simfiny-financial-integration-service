package database

import (
	"context"
	"fmt"
	"reflect"
	"sync"
	"testing"
	"time"

	postgresdb "github.com/SimifiniiCTO/simfiny-core-lib/database/postgres"
	"github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"
	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/dal"
	"go.uber.org/zap"
	"gorm.io/gorm"
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

func TestNew(t *testing.T) {
	type args struct {
		ctx  context.Context
		opts []Option
	}
	tests := []struct {
		name    string
		args    args
		want    *Db
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.ctx, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDb_Validate(t *testing.T) {
	tests := []struct {
		name    string
		db      *Db
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.db.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Db.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDb_performSchemaMigration(t *testing.T) {
	tests := []struct {
		name    string
		db      *Db
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.db.performSchemaMigration(); (err != nil) != tt.wantErr {
				t.Errorf("Db.performSchemaMigration() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDb_pingDatabase(t *testing.T) {
	tests := []struct {
		name    string
		db      *Db
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.db.pingDatabase(); (err != nil) != tt.wantErr {
				t.Errorf("Db.pingDatabase() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDb_configureConnection(t *testing.T) {
	tests := []struct {
		name    string
		db      *Db
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.db.configureConnection(); (err != nil) != tt.wantErr {
				t.Errorf("Db.configureConnection() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
