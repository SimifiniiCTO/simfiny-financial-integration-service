package grpc

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/helper"
	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/api/v1"
)

func TestServer_DeleteUserProfile(t *testing.T) {
	conn, client := setupPreconditions()
	defer conn.Close()

	// ensure the conn and client are not nil
	assert.NotNil(t, client)
	assert.NotNil(t, conn)

	type args struct {
		ctx context.Context
		req *proto.DeleteUserProfileRequest
	}
	tests := []struct {
		name                    string
		args                    args
		wantErr                 bool
		shouldCreateUserProfile bool
	}{
		// TODO: Add test cases.
		{
			name: "[success] DeleteUserProfile",
			args: args{
				ctx: context.Background(),
				req: &proto.DeleteUserProfileRequest{
					UserId: uint64(helper.GenerateRandomId(10000, 11000)),
				},
			},
			wantErr:                 false,
			shouldCreateUserProfile: true,
		},
		{
			name: "[failure] DeleteUserProfile - invalid request",
			args: args{
				ctx: context.Background(),
				req: &proto.DeleteUserProfileRequest{
					UserId: 0,
				},
			},
			wantErr:                 true,
			shouldCreateUserProfile: false,
		},
		{
			name: "[failure] DeleteUserProfile - user profile not found",
			args: args{
				ctx: context.Background(),
				req: &proto.DeleteUserProfileRequest{
					UserId: 1,
				},
			},
			wantErr:                 true,
			shouldCreateUserProfile: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldCreateUserProfile {
				_, err := client.CreateUserProfile(tt.args.ctx, &proto.CreateUserProfileRequest{
					Profile: &proto.UserProfile{
						UserId: tt.args.req.UserId,
					},
					Email: fmt.Sprintf("%s@gmail.com", helper.GenerateRandomString(10)),
				})

				assert.Nil(t, err)
			}

			got, err := client.DeleteUserProfile(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.DeleteUserProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && got == nil {
				t.Errorf("Server.DeleteUserProfile() got = %v, want %v", got, tt.wantErr)
			}
		})
	}
}
