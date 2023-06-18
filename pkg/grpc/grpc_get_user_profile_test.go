package grpc

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/helper"
	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/api/v1"
)

func TestServer_GetUserProfile(t *testing.T) {
	conn, client := setupPreconditions()
	defer conn.Close()

	// ensure the conn and client are not nil
	assert.NotNil(t, client)
	assert.NotNil(t, conn)

	type args struct {
		ctx     context.Context
		request *proto.GetUserProfileRequest
	}
	tests := []struct {
		name               string
		args               args
		wantErr            bool
		shouldCreateRecord bool
	}{
		{
			name: "[success] - get a valid user profile that already exists",
			args: args{
				ctx: context.Background(),
				request: &proto.GetUserProfileRequest{
					UserId: uint64(helper.GenerateRandomId(100000, 3000000)),
				},
			},
			wantErr:            false,
			shouldCreateRecord: true,
		},
		{
			name: "[failure] - get a user profile that does not exist",
			args: args{
				ctx: context.Background(),
				request: &proto.GetUserProfileRequest{
					UserId: uint64(helper.GenerateRandomId(100000, 3000000)),
				},
			},
			wantErr:            true,
			shouldCreateRecord: false,
		},
		{
			name: "[failure] - get a user profile with invalid userid",
			args: args{
				ctx: context.Background(),
				request: &proto.GetUserProfileRequest{
					UserId: 0,
				},
			},
			wantErr:            true,
			shouldCreateRecord: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldCreateRecord {
				_, err := client.CreateUserProfile(context.Background(), &proto.CreateUserProfileRequest{
					Profile: &proto.UserProfile{
						UserId: tt.args.request.UserId,
					},
					Email: fmt.Sprintf("%s@gmail.com", helper.GenerateRandomString(10)),
				})

				assert.Nil(t, err)
			}

			got, err := client.GetUserProfile(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.GetUserProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.NotNil(t, got)
				assert.Equal(t, tt.args.request.UserId, got.Profile.UserId)
			}
		})
	}
}
