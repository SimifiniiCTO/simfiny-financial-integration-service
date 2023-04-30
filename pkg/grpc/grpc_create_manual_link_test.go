package grpc

import (
	"context"
	"reflect"
	"testing"

	telemetry "github.com/SimifiniiCTO/core/core-telemetry"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/database"
	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/instrumentation"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/plaidhandler"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/secrets"
	transactionmanager "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/transaction_manager"
	"github.com/stripe/stripe-go/v74/client"
	"go.uber.org/zap"
)

func TestServer_CreateManualLink(t *testing.T) {
	type fields struct {
		UnimplementedFinancialServiceServer proto.UnimplementedFinancialServiceServer
		logger                              *zap.Logger
		config                              *Config
		instrumentation                     *instrumentation.ServiceTelemetry
		conn                                *database.Db
		plaidClient                         *plaidhandler.PlaidWrapper
		MetricsEngine                       *telemetry.MetricsEngine
		stripeClient                        *client.API
		kms                                 secrets.KeyManagement
		TransactionManager                  *transactionmanager.TransactionManager
	}
	type args struct {
		ctx context.Context
		req *proto.CreateManualLinkRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *proto.CreateManualLinkResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				UnimplementedFinancialServiceServer: tt.fields.UnimplementedFinancialServiceServer,
				logger:                              tt.fields.logger,
				config:                              tt.fields.config,
				instrumentation:                     tt.fields.instrumentation,
				conn:                                tt.fields.conn,
				plaidClient:                         tt.fields.plaidClient,
				MetricsEngine:                       tt.fields.MetricsEngine,
				stripeClient:                        tt.fields.stripeClient,
				kms:                                 tt.fields.kms,
				TransactionManager:                  tt.fields.TransactionManager,
			}
			got, err := s.CreateManualLink(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.CreateManualLink() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.CreateManualLink() = %v, want %v", got, tt.want)
			}
		})
	}
}
