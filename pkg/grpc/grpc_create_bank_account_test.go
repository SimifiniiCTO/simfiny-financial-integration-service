package grpc

import (
	"context"
	"fmt"
	"testing"

	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/helper"
	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"github.com/stretchr/testify/assert"
)

func TestServer_CreateBankAccount(t *testing.T) {
	conn, client := setupPreconditions()
	defer conn.Close()

	// ensure the conn and client are not nil
	assert.NotNil(t, client)
	assert.NotNil(t, conn)

	type args struct {
		ctx context.Context
		req *proto.CreateBankAccountRequest
	}
	tests := []struct {
		name                string
		args                args
		wantErr             bool
		shouldCreateProfile bool
	}{
		{
			name: "[success] - create bank account",
			args: args{
				ctx: context.Background(),
				req: &proto.CreateBankAccountRequest{
					BankAccount: helper.GenerateBankAccount(),
					UserId:      uint64(helper.GenerateRandomId(10000, 11000)),
				},
			},
			shouldCreateProfile: true,
			wantErr:             false,
		},
		{
			name: "[failure] - create bank account - invalid request",
			args: args{
				ctx: context.Background(),
				req: &proto.CreateBankAccountRequest{
					BankAccount: &proto.BankAccount{
						Name: "",
					},
					UserId: 0,
				},
			},
			shouldCreateProfile: false,
			wantErr:             true,
		},
		{

			name: "[failure] - create bank account - user profile not found",
			args: args{
				ctx: context.Background(),
				req: &proto.CreateBankAccountRequest{
					BankAccount: helper.GenerateBankAccount(),
					UserId:      1,
				},
			},
			shouldCreateProfile: false,
			wantErr:             true,
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

			got, err := client.CreateBankAccount(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.CreateBankAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.NotNil(t, got)
			}
		})
	}
}
