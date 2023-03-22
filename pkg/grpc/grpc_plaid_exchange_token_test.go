package grpc

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/helper"
)

func TestServer_PlaidExchangeToken(t *testing.T) {
	conn, client := setupPreconditions()
	defer conn.Close()

	// ensure the conn and client are not nil
	assert.NotNil(t, client)
	assert.NotNil(t, conn)

	type args struct {
		ctx context.Context
		req *proto.PlaidExchangeTokenRequest
	}
	tests := []struct {
		name                string
		args                args
		wantErr             bool
		shouldCreateProfile bool
	}{
		{
			name: "[success] - exchange token",
			args: args{
				ctx: context.Background(),
				req: &proto.PlaidExchangeTokenRequest{
					UserId:          uint64(helper.GenerateRandomId(10000, 11000)),
					PublicToken:     helper.GenerateRandomString(10),
					InstitutionId:   helper.GenerateRandomString(10),
					InstitutionName: helper.GenerateRandomString(10),
				},
			},
			wantErr:             false,
			shouldCreateProfile: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldCreateProfile {
				_, err := client.CreateUserProfile(context.Background(), &proto.CreateUserProfileRequest{
					Profile: &proto.UserProfile{
						UserId: tt.args.req.UserId,
					},
					Email: fmt.Sprintf("%s@gmail.com", helper.GenerateRandomString(10)),
				})
				assert.Nil(t, err)
			}

			got, err := client.PlaidExchangeToken(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.PlaidExchangeToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.NotNil(t, got)
			}
		})
	}
}
