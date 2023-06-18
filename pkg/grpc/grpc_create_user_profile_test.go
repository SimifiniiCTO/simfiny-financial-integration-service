package grpc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"

	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/helper"
	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
)

func TestServer_CreateUserProfile(t *testing.T) {
	conn, client := setupPreconditions()
	defer conn.Close()

	// ensure the conn and client are not nil
	assert.NotNil(t, client)
	assert.NotNil(t, conn)

	type args struct {
		ctx context.Context
		req *proto.CreateUserProfileRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "CreateUserProfile",
			args: args{
				ctx: context.Background(),
				req: &proto.CreateUserProfileRequest{
					Profile: &proto.UserProfile{
						UserId: uint64(helper.GenerateRandomId(10000, 11000)),
					},
					Email: "yoan@gmail.com",
				},
			},
			wantErr: false,
		},
		{
			name: "CreateUserProfile - invalid request",
			args: args{
				ctx: context.Background(),
				req: &proto.CreateUserProfileRequest{
					Profile: &proto.UserProfile{
						UserId: 0,
					},
					Email: "dhgjsdhg@gmail.com",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.CreateUserProfile(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.CreateUserProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.NotNil(t, got)
			}
		})
	}
}

func setupPreconditions() (*grpc.ClientConn, proto.FinancialServiceClient) {
	ctx := context.Background()
	conn := MockGRPCService(ctx)
	c := proto.NewFinancialServiceClient(conn)
	return conn, c
}
