package grpc

import (
	"context"
	"reflect"
	"testing"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

func TestServer_CreateManualLink(t *testing.T) {
	type args struct {
		ctx context.Context
		req *proto.CreateManualLinkRequest
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		want    *proto.CreateManualLinkResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.CreateManualLink(tt.args.ctx, tt.args.req)
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

func TestServer_autoGenerateManualBankAccount(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID *uint64
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		want    *proto.BankAccount
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.autoGenerateManualBankAccount(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.autoGenerateManualBankAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.autoGenerateManualBankAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}
