package grpc

import (
	"context"
	"reflect"
	"testing"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/api/v1"
)

func TestServer_DeleteBankAccount(t *testing.T) {
	type args struct {
		ctx context.Context
		req *proto.DeleteBankAccountRequest
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		want    *proto.DeleteBankAccountResponse
		wantErr bool
	}{
		// TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.DeleteBankAccount(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.DeleteBankAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.DeleteBankAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}
