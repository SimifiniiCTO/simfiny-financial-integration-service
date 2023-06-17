package grpc

import (
	"context"
	"reflect"
	"testing"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

func TestServer_GetLinks(t *testing.T) {
	type args struct {
		ctx context.Context
		req *proto.GetLinksRequest
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		want    *proto.GetLinksResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetLinks(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.GetLinks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.GetLinks() = %v, want %v", got, tt.want)
			}
		})
	}
}
