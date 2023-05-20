package grpc

import (
	"context"
	"fmt"
	"testing"
	"time"

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
			userId := tt.args.req.UserId
			if tt.shouldCreateProfile {
				res, err := client.CreateUserProfile(context.Background(), &proto.CreateUserProfileRequest{
					Profile: &proto.UserProfile{
						UserId: tt.args.req.UserId,
					},
					Email: fmt.Sprintf("%s@gmail.com", helper.GenerateRandomString(10)),
				})
				assert.Nil(t, err)

				userId = res.UserId
			}

			tt.args.req.UserId = userId

			got, err := client.PlaidExchangeToken(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.PlaidExchangeToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			time.Sleep(time.Second * 40)

			if !tt.wantErr {
				res, err := client.GetUserProfile(tt.args.ctx, &proto.GetUserProfileRequest{
					UserId: userId,
				})

				assert.Nil(t, err)
				assert.NotNil(t, res)

				txns, cnt, err := MockServer.clickhouseConn.GetTransactions(tt.args.ctx, &userId, int64(1), int64(100))
				assert.NoError(t, err)
				assert.NotNil(t, txns)
				assert.NotNil(t, cnt)

				assert.NotNil(t, got)
			}
		})
	}
}
