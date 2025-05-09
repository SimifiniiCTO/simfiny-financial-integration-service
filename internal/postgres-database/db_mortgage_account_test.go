package database

import (
	"context"
	"reflect"
	"testing"
	"time"

	postgresdb "github.com/SimifiniiCTO/simfiny-core-lib/database/postgres"
	"github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/secrets"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/dal"
	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"go.uber.org/zap"
)

func TestDb_GetMortgageAccount(t *testing.T) {
	type fields struct {
		Conn                   *postgresdb.Client
		Logger                 *zap.Logger
		MaxConnectionAttempts  int
		MaxRetriesPerOperation int
		RetryTimeOut           time.Duration
		OperationSleepInterval time.Duration
		Instrumentation        *instrumentation.Client
		QueryOperator          *dal.Query
		Kms                    secrets.KeyManagement
	}
	type args struct {
		ctx               context.Context
		MortgageAccountId uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *schema.MortgageAccount
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &Db{
				Conn:                   tt.fields.Conn,
				Logger:                 tt.fields.Logger,
				MaxConnectionAttempts:  tt.fields.MaxConnectionAttempts,
				MaxRetriesPerOperation: tt.fields.MaxRetriesPerOperation,
				RetryTimeOut:           tt.fields.RetryTimeOut,
				OperationSleepInterval: tt.fields.OperationSleepInterval,
				Instrumentation:        tt.fields.Instrumentation,
				QueryOperator:          tt.fields.QueryOperator,
				Kms:                    tt.fields.Kms,
			}
			got, err := db.GetMortgageAccount(tt.args.ctx, tt.args.MortgageAccountId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Db.GetMortgageAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Db.GetMortgageAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}
