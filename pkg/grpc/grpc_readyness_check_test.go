package grpc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

func TestServer_ReadynessCheck(t *testing.T) {
	conn, client := setupPreconditions()
	defer conn.Close()

	// ensure the conn and client are not nil
	assert.NotNil(t, client)
	assert.NotNil(t, conn)

	type args struct {
		ctx context.Context
		req *proto.ReadynessCheckRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test ReadynessCheck",
			args: args{
				ctx: context.Background(),
				req: &proto.ReadynessCheckRequest{},
			},
			wantErr: false,
		},
		{
			name: "Test ReadynessCheck - invalid request",
			args: args{
				ctx: context.Background(),
				req: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.ReadynessCheck(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.ReadynessCheck() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				assert.NotNil(t, got)
				assert.True(t, got.Healthy)
			}
		})
	}
}
