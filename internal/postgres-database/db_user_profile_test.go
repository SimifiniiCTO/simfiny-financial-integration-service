package database

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	postgresdb "github.com/SimifiniiCTO/simfiny-core-lib/database/postgres"
	"github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"
	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/dal"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/helper"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/secrets"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestDb_CreateUserProfile(t *testing.T) {
	var (
		validUserID          = uint64(helper.GenerateRandomId(10000, 3000000))
		invalidUserID uint64 = 0
	)

	type args struct {
		ctx     context.Context
		profile *schema.UserProfile
	}
	tests := []struct {
		name       string
		args       args
		wantUserId uint64
		wantErr    bool
	}{
		// TODO: Add test cases.
		{
			name: "[success] - create user profile",
			args: args{
				ctx: context.Background(),
				profile: &schema.UserProfile{
					UserId: validUserID,
				},
			},
			wantUserId: validUserID,
			wantErr:    false,
		},
		{
			name: "[failure] - create user profile with invalid user id",
			args: args{
				ctx: context.Background(),
				profile: &schema.UserProfile{
					UserId: invalidUserID,
				},
			},
			wantUserId: invalidUserID,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := conn.CreateUserProfile(tt.args.ctx, tt.args.profile)
			if (err != nil) != tt.wantErr {
				t.Errorf("Db.CreateUserProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if !reflect.DeepEqual(got.UserId, tt.wantUserId) {
					t.Errorf("Db.CreateUserProfile() = %v, want %v", got, tt.wantUserId)
				}
			}
		})
	}
}

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
					UserId: tt.args.userID,
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
				profile: helper.GenereateRandomUserProfileForTest(),
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

func TestDb_GetUserProfileByUserID(t *testing.T) {
	validUserProfile := helper.GenereateRandomUserProfileForTest()

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
				assert.NotNil(t, got.Link)
				assert.NotNil(t, got.Id)
			}
		})
	}
}

func TestDb_UpdateUserProfileSubscription(t *testing.T) {
	type fields struct {
		Conn                   *postgresdb.Client
		Logger                 *zap.Logger
		MaxConnectionAttempts  int
		MaxRetriesPerOperation int
		RetryTimeOut           time.Duration
		OperationSleepInterval time.Duration
		Instrumentation        *instrumentation.Client
		QueryOperator          *dal.Query
		Kms                    secrets.KeyManagement
	}
	type args struct {
		ctx          context.Context
		userId       uint64
		subscription *schema.StripeSubscription
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &Db{
				Conn:                   tt.fields.Conn,
				Logger:                 tt.fields.Logger,
				MaxConnectionAttempts:  tt.fields.MaxConnectionAttempts,
				MaxRetriesPerOperation: tt.fields.MaxRetriesPerOperation,
				RetryTimeOut:           tt.fields.RetryTimeOut,
				OperationSleepInterval: tt.fields.OperationSleepInterval,
				Instrumentation:        tt.fields.Instrumentation,
				QueryOperator:          tt.fields.QueryOperator,
				Kms:                    tt.fields.Kms,
			}
			if err := db.UpdateUserProfileSubscription(tt.args.ctx, tt.args.userId, tt.args.subscription); (err != nil) != tt.wantErr {
				t.Errorf("Db.UpdateUserProfileSubscription() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
