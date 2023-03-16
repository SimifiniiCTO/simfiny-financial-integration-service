package database

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/generated/api/v1"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/helper"
)

func TestDb_GetUserProfileByUserID(t *testing.T) {
	validUserProfile := genereateRandomUserProfileForTest()

	type args struct {
		ctx         context.Context
		userID      uint64
		userProfile *schema.UserProfile
	}
	tests := []struct {
		name               string
		description        string
		args               args
		shouldCreateRecord bool
		wantErr            bool
	}{
		{
			name:        "[success] - get a valid user",
			description: "get a valid user that is already existing",
			args: args{
				ctx:         context.Background(),
				userID:      validUserProfile.UserId,
				userProfile: validUserProfile,
			},
			wantErr:            false,
			shouldCreateRecord: true,
		},
		{
			name:        "[failure] - attempt to get a user record with invalid userid",
			description: "attempt to get a user record with invalid userid",
			args: args{
				ctx:         context.Background(),
				userID:      0,
				userProfile: nil,
			},
			wantErr:            true,
			shouldCreateRecord: false,
		},
		{
			name:        "[failure] - attempt to get a user record that does not exist",
			description: "attempt to get a user record that does not exist",
			args: args{
				ctx:         context.Background(),
				userID:      uint64(helper.GenerateRandomId(100000, 3000000)),
				userProfile: nil,
			},
			wantErr:            true,
			shouldCreateRecord: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldCreateRecord {
				record, err := conn.CreateUserProfile(tt.args.ctx, tt.args.userProfile)
				if err != nil {
					assert.Nil(t, err, fmt.Errorf("Db.CreateUserProfile() error = %v, wantErr %v", err, tt.wantErr))
				}

				// associate created record to the profile object
				tt.args.userProfile = record
			}

			got, err := conn.GetUserProfileByUserID(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Db.GetUserProfileByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.Equal(t, tt.args.userProfile.UserId, got.UserId)
				assert.NotNil(t, got.BankAccounts)
				assert.NotNil(t, got.CreditAccounts)
			}
		})
	}
}
