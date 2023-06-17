package grpc

import (
	"context"
	"reflect"
	"testing"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

func TestServer_DeleteLink(t *testing.T) {
	type args struct {
		ctx context.Context
		req *proto.DeleteLinkRequest
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		want    *proto.DeleteLinkResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.DeleteLink(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.DeleteLink() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.DeleteLink() = %v, want %v", got, tt.want)
			}
		})
	}
}
