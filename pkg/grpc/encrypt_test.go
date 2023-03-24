package grpc

import (
	"context"
	"reflect"
	"testing"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

func TestServer_EncryptAccessToken(t *testing.T) {
	type args struct {
		ctx         context.Context
		accessToken string
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		want    *accessTokenMeta
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.EncryptAccessToken(tt.args.ctx, tt.args.accessToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.EncryptAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.EncryptAccessToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_DecryptUserAccessToken(t *testing.T) {
	type args struct {
		ctx   context.Context
		token *proto.Token
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		want    *string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.DecryptUserAccessToken(tt.args.ctx, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.DecryptUserAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.DecryptUserAccessToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
