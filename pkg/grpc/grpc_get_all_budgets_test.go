package grpc

import (
	"context"
	"reflect"
	"testing"

	telemetry "github.com/SimifiniiCTO/core/core-telemetry"
	"github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/plaidhandler"
	database "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/postgres-database"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/secrets"
	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"github.com/stripe/stripe-go/v74/client"
	"go.uber.org/zap"
)

func TestServer_getPocketBudgets(t *testing.T) {
	type fields struct {
		UnimplementedFinancialServiceServer proto.UnimplementedFinancialServiceServer
		logger                              *zap.Logger
		config                              *Config
		instrumentation                     *instrumentation.Client
		conn                                *database.Db
		plaidClient                         *plaidhandler.PlaidWrapper
		MetricsEngine                       *telemetry.MetricsEngine
		stripeClient                        *client.API
		kms                                 secrets.KeyManagement
	}
	type args struct {
		pocket *proto.Pocket
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*proto.Budget
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
			}
			if got := s.getPocketBudgets(tt.args.pocket); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.getPocketBudgets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_GetAllBudgets(t *testing.T) {
	type args struct {
		ctx context.Context
		req *proto.GetAllBudgetsRequest
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		want    *proto.GetAllBudgetsResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetAllBudgets(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.GetAllBudgets() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.GetAllBudgets() = %v, want %v", got, tt.want)
			}
		})
	}
}
