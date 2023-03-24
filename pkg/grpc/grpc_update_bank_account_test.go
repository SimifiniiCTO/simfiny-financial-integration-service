package grpc

import (
	"context"
	"reflect"
	"testing"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

func TestServer_UpdateBankAccount(t *testing.T) {
	type args struct {
		in0 context.Context
		in1 *proto.UpdateBankAccountRequest
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		want    *proto.UpdateBankAccountResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.UpdateBankAccount(tt.args.in0, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.UpdateBankAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.UpdateBankAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}
