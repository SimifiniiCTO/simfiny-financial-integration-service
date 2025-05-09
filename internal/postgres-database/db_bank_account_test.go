package database

import (
	"context"
	"testing"

	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/helper"
	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"github.com/stretchr/testify/assert"
)

func TestDb_CreateBankAccount(t *testing.T) {
	type args struct {
		ctx         context.Context
		userID      uint64
		linkID      uint64
		bankAccount *schema.BankAccount
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
				ctx:         context.Background(),
				userID:      uint64(helper.GenerateRandomId(10000, 3000000)),
				linkID:      uint64(helper.GenerateRandomId(10000, 3000000)),
				bankAccount: generateBankAccount(),
			},
			shouldCreateProfile: true,
			wantErr:             false,
		},
		{
			name: "[failure] - create bank account with invalid user id",
			args: args{
				ctx:         context.Background(),
				userID:      0,
				linkID:      uint64(helper.GenerateRandomId(10000, 3000000)),
				bankAccount: generateBankAccount(),
			},
			wantErr:             true,
			shouldCreateProfile: false,
		},
		{
			name: "[failure] - create bank account with invalid bank account",
			args: args{
				ctx:         context.Background(),
				userID:      uint64(helper.GenerateRandomId(10000, 3000000)),
				linkID:      uint64(helper.GenerateRandomId(10000, 3000000)),
				bankAccount: nil,
			},
			wantErr:             true,
			shouldCreateProfile: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldCreateProfile {
				_, err := conn.CreateUserProfile(tt.args.ctx, &schema.UserProfile{
					UserId: tt.args.userID,
				})
				assert.Nil(t, err)

				// create a link for the given user
				link, err := conn.CreateLink(tt.args.ctx, tt.args.userID, helper.GenerateLink(schema.LinkType_LINK_TYPE_PLAID), false)
				assert.Nil(t, err)

				tt.args.linkID = link.Id
			}

			got, err := conn.CreateBankAccount(tt.args.ctx, tt.args.linkID, tt.args.bankAccount)
			if (err != nil) != tt.wantErr {
				t.Errorf("Db.CreateBankAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.NotNil(t, got)
				assert.Greater(t, got.Id, uint64(0))
			}
		})
	}
}

func TestDb_DeleteBankAccount(t *testing.T) {
	type args struct {
		ctx           context.Context
		bankAccountID uint64
		userID        uint64
		bankAccount   *schema.BankAccount
	}
	tests := []struct {
		name                string
		args                args
		wantErr             bool
		shouldCreateProfile bool
	}{
		{
			name: "[success] - delete a valid bank account",
			args: args{
				ctx:           context.Background(),
				userID:        uint64(helper.GenerateRandomId(10000, 3000000)),
				bankAccount:   generateBankAccount(),
				bankAccountID: uint64(helper.GenerateRandomId(10000, 3000000)),
			},
			shouldCreateProfile: true,
			wantErr:             false,
		},
		{
			name: "[failure] - delete a bank account that does not exist",
			args: args{
				ctx:           context.Background(),
				userID:        0,
				bankAccount:   generateBankAccount(),
				bankAccountID: uint64(helper.GenerateRandomId(10000, 3000000)),
			},
			wantErr:             true,
			shouldCreateProfile: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldCreateProfile {
				_, err := conn.CreateUserProfile(tt.args.ctx, &schema.UserProfile{
					UserId: tt.args.userID,
				})
				assert.Nil(t, err)

				// create a link for the given user
				link, err := conn.CreateLink(tt.args.ctx, tt.args.userID, helper.GenerateLink(schema.LinkType_LINK_TYPE_PLAID), false)
				assert.Nil(t, err)

				newAcct, err := conn.CreateBankAccount(tt.args.ctx, link.Id, tt.args.bankAccount)
				assert.Nil(t, err)

				tt.args.bankAccountID = newAcct.Id
			}

			if err := conn.DeleteBankAccount(tt.args.ctx, tt.args.bankAccountID); (err != nil) != tt.wantErr {
				t.Errorf("Db.DeleteBankAccount() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				_, err := conn.GetBankAccount(tt.args.ctx, tt.args.bankAccountID)
				assert.NotNil(t, err)
			}
		})
	}
}

func TestDb_GetBankAccount(t *testing.T) {
	type args struct {
		ctx           context.Context
		bankAccountID uint64
		userID        uint64
		bankAccount   *schema.BankAccount
	}
	tests := []struct {
		name                string
		args                args
		shouldCreateProfile bool
		wantErr             bool
	}{
		{
			name: "[success] - get a bank account that exists",
			args: args{
				ctx:           context.Background(),
				userID:        uint64(helper.GenerateRandomId(10000, 3000000)),
				bankAccount:   generateBankAccount(),
				bankAccountID: uint64(helper.GenerateRandomId(10000, 3000000)),
			},
			shouldCreateProfile: true,
		},
		{
			name: "[failure] - get a bank account that does not exist",
			args: args{
				ctx:           context.Background(),
				userID:        0,
				bankAccount:   generateBankAccount(),
				bankAccountID: uint64(helper.GenerateRandomId(10000, 3000000)),
			},
			wantErr:             true,
			shouldCreateProfile: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldCreateProfile {
				_, err := conn.CreateUserProfile(tt.args.ctx, &schema.UserProfile{
					UserId: tt.args.userID,
				})
				assert.Nil(t, err)

				// create a link for the given user
				link, err := conn.CreateLink(tt.args.ctx, tt.args.userID, helper.GenerateLink(schema.LinkType_LINK_TYPE_PLAID), false)
				assert.Nil(t, err)

				acct, err := conn.CreateBankAccount(tt.args.ctx, link.Id, tt.args.bankAccount)
				assert.Nil(t, err)

				tt.args.bankAccountID = acct.Id
			}

			got, err := conn.GetBankAccount(tt.args.ctx, tt.args.bankAccountID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Db.GetBankAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.NotNil(t, got)
				assert.Greater(t, got.Id, uint64(0))
			}
		})
	}
}

func TestDb_UpdateBankAccount(t *testing.T) {
	type args struct {
		ctx           context.Context
		bankAccountID uint64
		userID        uint64
		bankAccount   *schema.BankAccount
	}
	tests := []struct {
		name                string
		args                args
		shouldCreateProfile bool
		wantErr             bool
	}{
		{
			name: "[success] - get a bank account that exists",
			args: args{
				ctx:           context.Background(),
				userID:        uint64(helper.GenerateRandomId(10000, 3000000)),
				bankAccount:   generateBankAccount(),
				bankAccountID: uint64(helper.GenerateRandomId(10000, 3000000)),
			},
			shouldCreateProfile: true,
		},
		{
			name: "[failure] - get a bank account that does not exist",
			args: args{
				ctx:           context.Background(),
				userID:        0,
				bankAccount:   generateBankAccount(),
				bankAccountID: uint64(helper.GenerateRandomId(10000, 3000000)),
			},
			wantErr:             true,
			shouldCreateProfile: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updateBankAcctName := "new name"
			if tt.shouldCreateProfile {
				_, err := conn.CreateUserProfile(tt.args.ctx, &schema.UserProfile{
					UserId: tt.args.userID,
				})
				assert.Nil(t, err)

				// create a link for the given user
				link, err := conn.CreateLink(tt.args.ctx, tt.args.userID, helper.GenerateLink(schema.LinkType_LINK_TYPE_PLAID), false)
				assert.Nil(t, err)

				acct, err := conn.CreateBankAccount(tt.args.ctx, link.Id, tt.args.bankAccount)
				assert.Nil(t, err)

				// change the name
				acct.Name = updateBankAcctName
				tt.args.bankAccount = acct
			}

			if err := conn.UpdateBankAccount(tt.args.ctx, tt.args.bankAccount); (err != nil) != tt.wantErr {
				t.Errorf("Db.UpdateBankAccount() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				got, err := conn.GetBankAccount(tt.args.ctx, tt.args.bankAccount.Id)
				assert.Nil(t, err)
				assert.NotNil(t, got)
				assert.Equal(t, got.Name, updateBankAcctName)
			}
		})
	}
}

func TestDb_UpdateBankAccounts(t *testing.T) {
	type args struct {
		ctx          context.Context
		bankAccounts []*schema.BankAccount
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
			if err := tt.db.UpdateBankAccounts(tt.args.ctx, nil, tt.args.bankAccounts); (err != nil) != tt.wantErr {
				t.Errorf("Db.UpdateBankAccounts() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
