package database

import (
	"context"
	"reflect"
	"testing"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/generated/api/v1"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/helper"
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
