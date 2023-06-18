package grpc

import (
	"context"
	"reflect"
	"testing"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
)

func TestServer_GetReOcurringTransactions(t *testing.T) {
	type args struct {
		ctx context.Context
		req *proto.GetReCurringTransactionsRequest
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		want    *proto.GetReCurringTransactionsResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetReOcurringTransactions(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.GetReOcurringTransactions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.GetReOcurringTransactions() = %v, want %v", got, tt.want)
			}
		})
	}
}
