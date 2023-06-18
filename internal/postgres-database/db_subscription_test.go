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

func TestDb_GetSubscriptionBySubscriptionId(t *testing.T) {

	type args struct {
		ctx            context.Context
		subscriptionId *string
	}
	tests := []struct {
		name    string
		args    args
		want    *schema.StripeSubscription
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := conn.GetSubscriptionBySubscriptionId(tt.args.ctx, tt.args.subscriptionId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Db.GetSubscriptionBySubscriptionId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Db.GetSubscriptionBySubscriptionId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDb_UpdateSubscription(t *testing.T) {
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
		ctx                  context.Context
		stripeSubscriptionId *string
		subscription         *schema.StripeSubscription
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
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
			if err := db.UpdateSubscription(tt.args.ctx, tt.args.stripeSubscriptionId, tt.args.subscription); (err != nil) != tt.wantErr {
				t.Errorf("Db.UpdateSubscription() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
