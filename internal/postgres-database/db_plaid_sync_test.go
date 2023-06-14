package database

import (
	"context"
	"testing"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/helper"
	"github.com/stretchr/testify/assert"
)

func TestDb_GetLastPlaidSync(t *testing.T) {
	type args struct {
		ctx          context.Context
		precondition func(ctx context.Context, t *testing.T) (uint64, uint64)
	}
	tests := []struct {
		name         string
		args         args
		wantErr      bool
		shouldCreate bool
	}{
		{
			name: "[success] - get last plaid sync",
			args: args{
				ctx: context.Background(),
				precondition: func(ctx context.Context, t *testing.T) (uint64, uint64) {
					userId := uint64(helper.GenerateRandomId(10000, 3000000))

					// create a user account
					_, err := conn.CreateUserProfile(ctx, &schema.UserProfile{
						UserId: userId,
					})
					assert.Nil(t, err)

					// create a "plaid" link for the given user
					link, err := conn.CreateLink(ctx, userId, helper.GenerateLink(schema.LinkType_LINK_TYPE_PLAID))
					assert.Nil(t, err)

					err = conn.RecordPlaidSync(ctx, userId, link.GetId(), "test", "", 0, 0, 0)
					assert.NoError(t, err)
					return userId, link.Id
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userId, linkId := tt.args.precondition(tt.args.ctx, t)
			got, err := conn.GetLastPlaidSync(tt.args.ctx, userId, linkId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Db.GetLastPlaidSync() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.NotNil(t, got)
			}
		})
	}
}

func TestDb_RecordPlaidSync(t *testing.T) {
	type args struct {
		ctx         context.Context
		userId      uint64
		plaidLinkId uint64
		trigger     string
		nextCursor  string
		added       int64
		modified    int64
		removed     int64
	}
	tests := []struct {
		name    string
		db      *Db
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.db.RecordPlaidSync(tt.args.ctx, tt.args.userId, tt.args.plaidLinkId, tt.args.trigger, tt.args.nextCursor, tt.args.added, tt.args.modified, tt.args.removed); (err != nil) != tt.wantErr {
				t.Errorf("Db.RecordPlaidSync() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
