package grpc

import (
	"context"
	"reflect"
	"testing"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
)

func TestServer_GetLink(t *testing.T) {
	type args struct {
		ctx context.Context
		req *proto.GetLinkRequest
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		want    *proto.GetLinkResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetLink(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.GetLink() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.GetLink() = %v, want %v", got, tt.want)
			}
		})
	}
}
