package database

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/helper"
)

func TestDb_DeleteUserProfileByUserID(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID uint64
	}
	tests := []struct {
		name               string
		description        string
		args               args
		wantErr            bool
		shouldCreateRecord bool
	}{
		// TODO: Add test cases.
		{
			name:        "[success] - delete a valid user",
			description: "delete a valid user that is already existing",
			args: args{
				ctx:    context.Background(),
				userID: uint64(helper.GenerateRandomId(100000, 3000000)),
			},
			wantErr:            false,
			shouldCreateRecord: true,
		},
		{
			name:        "[failure] - attempt to delete a user record with invalid userid",
			description: "attempt to delete a user record with invalid userid",
			args: args{
				ctx:    context.Background(),
				userID: 0,
			},
			wantErr:            true,
			shouldCreateRecord: false,
		},
		{
			name:        "[failure] - attempt to delete a user record that does not exist",
			description: "attempt to delete a user record that does not exist",
			args: args{
				ctx:    context.Background(),
				userID: uint64(helper.GenerateRandomId(100000, 3000000)),
			},
			wantErr:            true,
			shouldCreateRecord: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldCreateRecord {
				newprofile := &schema.UserProfile{
					UserId:           tt.args.userID,
					PlaidAccessToken: helper.GenerateRandomString(30),
				}

				_, err := conn.CreateUserProfile(tt.args.ctx, newprofile)
				if err != nil {
					assert.Nil(t, err, fmt.Errorf("Db.CreateUserProfile() error = %v, wantErr %v", err, tt.wantErr))
				}
			}

			if err := conn.DeleteUserProfileByUserID(tt.args.ctx, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("Db.DeleteUserProfileByUserID() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				_, err := conn.GetUserProfileByUserID(tt.args.ctx, tt.args.userID)
				assert.NotNil(t, err)
			}
		})
	}
}
