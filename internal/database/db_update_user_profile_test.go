package database

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/helper"
)

func TestDb_UpdateUserProfile(t *testing.T) {
	type args struct {
		ctx     context.Context
		profile *schema.UserProfile
	}
	tests := []struct {
		name        string
		description string
		args        args
		// wether or not to create the record
		// prior to performing an update operation
		shouldCreateRecord bool
		wantErr            bool
	}{
		{
			name:        "[success] - update a valid user",
			description: "update a valid user that is already existing",
			args: args{
				ctx:     context.Background(),
				profile: genereateRandomUserProfileForTest(),
			},
			shouldCreateRecord: true,
			wantErr:            false,
		},
		{
			name:        "[failure] - update an invalid user record ",
			description: "update a user that is invalid",
			args: args{
				ctx:     context.Background(),
				profile: nil,
			},
			wantErr:            true,
			shouldCreateRecord: false,
		},
		{
			name:        "[failure] - update a user that does not exist",
			description: "update a user that does not exist",
			args: args{
				ctx: context.Background(),
				profile: &schema.UserProfile{
					UserId: uint64(helper.GenerateRandomId(100000, 3000000)),
				},
			},
			wantErr:            true,
			shouldCreateRecord: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldCreateRecord {
				record, err := conn.CreateUserProfile(tt.args.ctx, tt.args.profile)
				if err != nil {
					assert.Nil(t, err, fmt.Errorf("Db.CreateUserProfile() error = %v, wantErr %v", err, tt.wantErr))
				}

				// associate created record to the profile object
				tt.args.profile = record
			}

			// peform update to record
			// tt.args.profile.UserId = uint64(helper.GenerateRandomId(100000, 3000000))
			err := conn.UpdateUserProfile(tt.args.ctx, tt.args.profile)
			if (err != nil) != tt.wantErr {
				t.Errorf("Db.UpdateUserProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
