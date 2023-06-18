package database

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/helper"
	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/api/v1"
)

func TestDb_CreatePocket(t *testing.T) {
	type args struct {
		ctx           context.Context
		bankAccountID uint64
		pocket        *schema.Pocket
	}
	tests := []struct {
		name                    string
		args                    args
		wantErr                 bool
		shouldCreateBankAccount bool
	}{
		{
			name: "Create pocket",
			args: args{
				ctx:           context.Background(),
				bankAccountID: uint64(helper.GenerateRandomId(10000, 50000)),
				pocket:        helper.GenerateSinglePocket(),
			},
			wantErr:                 false,
			shouldCreateBankAccount: true,
		},
		{
			name: "Create pocket - BankAccountID not found",
			args: args{
				ctx:           context.Background(),
				bankAccountID: uint64(helper.GenerateRandomId(10000, 50000)),
				pocket:        helper.GenerateSinglePocket(),
			},
			wantErr:                 true,
			shouldCreateBankAccount: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldCreateBankAccount {
				_, bankAcct := createUserProfileAndBankAccountForTest(t, tt.args.ctx, conn)
				tt.args.bankAccountID = bankAcct.Id
			}

			got, err := conn.CreatePocket(tt.args.ctx, tt.args.bankAccountID, tt.args.pocket)
			if (err != nil) != tt.wantErr {
				t.Errorf("Db.CreatePocket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.NotNil(t, got)
			}
		})
	}
}

func TestDb_DeletePocket(t *testing.T) {
	type args struct {
		ctx           context.Context
		bankAccountID uint64
		pocketID      uint64
	}
	tests := []struct {
		name                    string
		args                    args
		wantErr                 bool
		shouldCreateBankAccount bool
	}{
		{
			name: "Delete pocket",
			args: args{
				ctx:           context.Background(),
				bankAccountID: uint64(helper.GenerateRandomId(10000, 50000)),
				pocketID:      uint64(helper.GenerateRandomId(10000, 50000)),
			},
			wantErr:                 false,
			shouldCreateBankAccount: true,
		},
		{
			name: "Delete pocket - BankAccountID not found",
			args: args{
				ctx:           context.Background(),
				bankAccountID: uint64(helper.GenerateRandomId(10000, 50000)),
				pocketID:      uint64(helper.GenerateRandomId(10000, 50000)),
			},
			wantErr:                 true,
			shouldCreateBankAccount: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldCreateBankAccount {
				_, bankAcct := createUserProfileAndBankAccountForTest(t, tt.args.ctx, conn)
				tt.args.bankAccountID = bankAcct.Id
				tt.args.pocketID = bankAcct.Pockets[0].Id
			}

			if err := conn.DeletePocket(tt.args.ctx, tt.args.pocketID); (err != nil) != tt.wantErr {
				t.Errorf("Db.DeletePocket() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				_, err := conn.GetPocket(tt.args.ctx, tt.args.pocketID)
				assert.NotNil(t, err)
			}
		})
	}
}

func TestDb_GetPocket(t *testing.T) {
	type args struct {
		ctx           context.Context
		bankAccountID uint64
		pocketID      uint64
	}
	tests := []struct {
		name                    string
		args                    args
		wantErr                 bool
		shouldCreateBankAccount bool
	}{
		{
			name: "Get pocket",
			args: args{
				ctx:           context.Background(),
				bankAccountID: uint64(helper.GenerateRandomId(10000, 50000)),
				pocketID:      uint64(helper.GenerateRandomId(10000, 50000)),
			},
			wantErr:                 false,
			shouldCreateBankAccount: true,
		},
		{
			name: "Get pocket - BankAccountID not found",
			args: args{
				ctx:           context.Background(),
				bankAccountID: uint64(helper.GenerateRandomId(10000, 50000)),
				pocketID:      uint64(helper.GenerateRandomId(10000, 50000)),
			},
			wantErr:                 true,
			shouldCreateBankAccount: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldCreateBankAccount {
				_, bankAcct := createUserProfileAndBankAccountForTest(t, tt.args.ctx, conn)
				tt.args.bankAccountID = bankAcct.Id
				tt.args.pocketID = bankAcct.Pockets[0].Id
			}

			got, err := conn.GetPocket(tt.args.ctx, tt.args.pocketID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Db.GetPocket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.NotNil(t, got)
			}
		})
	}
}

func TestDb_UpdatePocket(t *testing.T) {
	type args struct {
		ctx           context.Context
		bankAccountID uint64
		pocket        *schema.Pocket
	}
	tests := []struct {
		name                    string
		args                    args
		wantErr                 bool
		shouldCreateBankAccount bool
	}{
		{
			name: "Update pocket",
			args: args{
				ctx:           context.Background(),
				bankAccountID: uint64(helper.GenerateRandomId(10000, 50000)),
				pocket:        helper.GenerateSinglePocket(),
			},
			wantErr:                 false,
			shouldCreateBankAccount: true,
		},
		{
			name: "Update pocket - BankAccountID not found",
			args: args{
				ctx:           context.Background(),
				bankAccountID: uint64(helper.GenerateRandomId(10000, 50000)),
				pocket:        helper.GenerateSinglePocket(),
			},
			wantErr:                 true,
			shouldCreateBankAccount: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			const updatePocketType = schema.PocketType_POCKET_TYPE_FUN_MONEY

			if tt.shouldCreateBankAccount {
				_, bankAcct := createUserProfileAndBankAccountForTest(t, tt.args.ctx, conn)
				tt.args.bankAccountID = bankAcct.Id
				tt.args.pocket = bankAcct.Pockets[0]
				tt.args.pocket.Type = updatePocketType
			}

			if err := conn.UpdatePocket(tt.args.ctx, tt.args.pocket); (err != nil) != tt.wantErr {
				t.Errorf("Db.UpdatePocket() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				got, err := conn.GetPocket(tt.args.ctx, tt.args.pocket.Id)
				assert.Nil(t, err)
				assert.NotNil(t, got)
				assert.Equal(t, updatePocketType, got.Type)
			}
		})
	}
}
