package clickhousedatabase

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/SimifiniiCTO/simfiny-core-lib/database/clickhouse"
	"github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"
	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/dal"
	"github.com/newrelic/go-agent/v3/newrelic"
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
		queryOperator:         dal.Use(client.Engine),
		logger:                zap.NewNop(),
		instrumentationClient: &instrumentation.Client{},
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

func TestDb_startDatastoreSpan(t *testing.T) {
	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name string
		db   *Db
		args args
		want *newrelic.DatastoreSegment
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.db.startDatastoreSpan(tt.args.ctx, tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Db.startDatastoreSpan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMockInMemoryClickhouseDB(t *testing.T) {
	tests := []struct {
		name string
		want *Db
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMockInMemoryClickhouseDB(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMockInMemoryClickhouseDB() = %v, want %v", got, tt.want)
			}
		})
	}
}
