package database

import (
	"context"
	"testing"

	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/helper"
	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
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
					link, err := conn.CreateLink(ctx, userId, helper.GenerateLink(schema.LinkType_LINK_TYPE_PLAID), false)
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
	type funcArgs struct {
		userId      uint64
		plaidLinkId uint64
		trigger     string
		nextCursor  string
		added       int64
		modified    int64
		removed     int64
	}

	type args struct {
		ctx          context.Context
		precondition func(ctx context.Context, t *testing.T) funcArgs
	}
	tests := []struct {
		name    string
		db      *Db
		args    args
		wantErr bool
	}{
		{
			name: "[success] - record plaid sync",
			db:   conn,
			args: args{
				ctx: context.Background(),
				precondition: func(ctx context.Context, t *testing.T) funcArgs {
					userId := uint64(helper.GenerateRandomId(10000, 3000000))

					// create a user account
					_, err := conn.CreateUserProfile(ctx, &schema.UserProfile{
						UserId: userId,
					})
					assert.Nil(t, err)

					// create a "plaid" link for the given user
					link, err := conn.CreateLink(ctx, userId, helper.GenerateLink(schema.LinkType_LINK_TYPE_PLAID), false)
					assert.Nil(t, err)

					return funcArgs{
						userId:      userId,
						plaidLinkId: link.Id,
						trigger:     "test",
						nextCursor:  "",
						added:       int64(helper.GenerateRandomId(100, 1000)),
						modified:    int64(helper.GenerateRandomId(100, 1000)),
						removed:     int64(helper.GenerateRandomId(100, 1000)),
					}
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.args.precondition(tt.args.ctx, t)
			if err := tt.db.RecordPlaidSync(tt.args.ctx, res.userId, res.plaidLinkId, res.trigger, res.nextCursor, res.added, res.modified, res.removed); (err != nil) != tt.wantErr {
				t.Errorf("Db.RecordPlaidSync() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
