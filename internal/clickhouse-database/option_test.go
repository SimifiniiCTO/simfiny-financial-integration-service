package clickhousedatabase

import (
	"reflect"
	"testing"

	clickhousedb "github.com/SimifiniiCTO/simfiny-core-lib/database/clickhouse"
	"github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/dal"
	"go.uber.org/zap"
)

func TestWithDatabaseInstrumentation(t *testing.T) {
	type args struct {
		instrumentation *instrumentation.Client
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithDatabaseInstrumentation(tt.args.instrumentation); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithDatabaseInstrumentation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithDatabaseQueryOperator(t *testing.T) {
	type args struct {
		queryOperator *dal.Query
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithDatabaseQueryOperator(tt.args.queryOperator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithDatabaseQueryOperator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithDatabaseLogger(t *testing.T) {
	type args struct {
		logger *zap.Logger
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithDatabaseLogger(tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithDatabaseLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithDatabaseClient(t *testing.T) {
	type args struct {
		client *clickhousedb.Client
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithDatabaseClient(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithDatabaseClient() = %v, want %v", got, tt.want)
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
