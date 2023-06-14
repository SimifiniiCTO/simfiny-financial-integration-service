package database

import (
	"reflect"
	"testing"
	"time"

	postgresdb "github.com/SimifiniiCTO/simfiny-core-lib/database/postgres"
	"github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/dal"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/secrets"
	"go.uber.org/zap"
)

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

func TestWithDatabaseMaxConnectionAttempts(t *testing.T) {
	type args struct {
		maxConnectionAttempts int
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
			if got := WithDatabaseMaxConnectionAttempts(tt.args.maxConnectionAttempts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithDatabaseMaxConnectionAttempts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithDatabaseMaxRetriesPerOperation(t *testing.T) {
	type args struct {
		maxRetriesPerOperation int
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
			if got := WithDatabaseMaxRetriesPerOperation(tt.args.maxRetriesPerOperation); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithDatabaseMaxRetriesPerOperation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithDatabaseRetryTimeOut(t *testing.T) {
	type args struct {
		retryTimeOut time.Duration
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
			if got := WithDatabaseRetryTimeOut(tt.args.retryTimeOut); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithDatabaseRetryTimeOut() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithDatabaseOperationSleepInterval(t *testing.T) {
	type args struct {
		operationSleepInterval time.Duration
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
			if got := WithDatabaseOperationSleepInterval(tt.args.operationSleepInterval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithDatabaseOperationSleepInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}

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

func TestWithConnection(t *testing.T) {
	type args struct {
		params *postgresdb.Client
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
			if got := WithConnection(tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithConnection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithKms(t *testing.T) {
	type args struct {
		kms secrets.KeyManagement
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
			if got := WithKms(tt.args.kms); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithKms() = %v, want %v", got, tt.want)
			}
		})
	}
}
