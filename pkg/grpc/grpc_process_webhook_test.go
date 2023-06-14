package grpc

import (
	"context"
	"reflect"
	"testing"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

func TestPlaidClaims_Valid(t *testing.T) {
	tests := []struct {
		name    string
		c       PlaidClaims
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Valid(); (err != nil) != tt.wantErr {
				t.Errorf("PlaidClaims.Valid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServer_ProcessWebhook(t *testing.T) {
	type args struct {
		ctx context.Context
		req *proto.ProcessWebhookRequest
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		want    *proto.ProcessWebhookResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ProcessWebhook(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.ProcessWebhook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.ProcessWebhook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_processWebhook(t *testing.T) {
	type args struct {
		ctx context.Context
		req *proto.ProcessWebhookRequest
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.processWebhook(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("Server.processWebhook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
