package database

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/helper"
	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/api/v1"
)

func TestDb_CreateBudget(t *testing.T) {

	type args struct {
		ctx         context.Context
		pocketId    uint64
		milestoneId uint64
		budget      *schema.Budget
	}
	tests := []struct {
		name                    string
		args                    args
		wantErr                 bool
		shouldCreateBankAccount bool
	}{
		{
			name: "[success] - create budget for an existing milestone",
			args: args{
				ctx:         context.Background(),
				pocketId:    uint64(helper.GenerateRandomId(10000, 300000)),
				milestoneId: uint64(helper.GenerateRandomId(10000, 300000)),
				budget:      helper.GenerateSingleBudget(),
			},
			shouldCreateBankAccount: true,
			wantErr:                 false,
		},
		{
			name: "[failure] - create budget for an invalid milestone",
			args: args{
				ctx:         context.Background(),
				pocketId:    uint64(helper.GenerateRandomId(10000, 300000)),
				milestoneId: 0,
				budget:      helper.GenerateSingleBudget(),
			},
			wantErr:                 true,
			shouldCreateBankAccount: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldCreateBankAccount {
				sampleUserID := uint64(helper.GenerateRandomId(10000, 3000000))
				_, err := conn.CreateUserProfile(tt.args.ctx, &schema.UserProfile{
					UserId: sampleUserID,
				})

				assert.Nil(t, err)

				sampleBankAcct := helper.GenerateBankAccount()
				newAcct, err := conn.CreateBankAccount(tt.args.ctx, sampleUserID, sampleBankAcct)
				if err != nil {
					t.Errorf("Db.CreateBankAccount() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				// get a pocket from the acct
				// NOTE: this is guaranteed to work because we just created the bank account
				// which by default should have all the fields populated
				tt.args.pocketId = newAcct.Pockets[0].Id
				tt.args.milestoneId = newAcct.Pockets[0].Goals[0].Milestones[0].Id
			}

			got, err := conn.CreateBudget(tt.args.ctx, tt.args.milestoneId, tt.args.budget)
			if (err != nil) != tt.wantErr {
				t.Errorf("Db.CreateBudget() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.NotNil(t, got)
				assert.Greater(t, got.Id, uint64(0))
			}
		})
	}
}

func TestDb_DeleteBudget(t *testing.T) {
	type args struct {
		ctx      context.Context
		budgetID uint64
	}
	tests := []struct {
		name                    string
		args                    args
		wantErr                 bool
		shouldCreateBankAccount bool
	}{
		{
			name: "[success] - delete budget for an existing milestone",
			args: args{
				ctx:      context.Background(),
				budgetID: uint64(helper.GenerateRandomId(10000, 300000)),
			},
			shouldCreateBankAccount: true,
			wantErr:                 false,
		},
		{
			name: "[failure] - delete budget for an invalid milestone",
			args: args{
				ctx:      context.Background(),
				budgetID: uint64(helper.GenerateRandomId(10000, 300000)),
			},
			wantErr:                 true,
			shouldCreateBankAccount: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldCreateBankAccount {
				sampleUserID := uint64(helper.GenerateRandomId(10000, 3000000))
				_, err := conn.CreateUserProfile(tt.args.ctx, &schema.UserProfile{
					UserId: sampleUserID,
				})

				assert.Nil(t, err)

				sampleBankAcct := helper.GenerateBankAccount()
				newAcct, err := conn.CreateBankAccount(tt.args.ctx, sampleUserID, sampleBankAcct)
				if err != nil {
					t.Errorf("Db.CreateBankAccount() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				// get a pocket from the acct
				// NOTE: this is guaranteed to work because we just created the bank account
				// which by default should have all the fields populated
				tt.args.budgetID = newAcct.Pockets[0].Goals[0].Milestones[0].Budget.Id
			}

			if err := conn.DeleteBudget(tt.args.ctx, tt.args.budgetID); (err != nil) != tt.wantErr {
				t.Errorf("Db.DeleteBudget() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				// check if the budget was deleted
				_, err := conn.GetBudget(tt.args.ctx, tt.args.budgetID)
				assert.NotNil(t, err)
			}
		})
	}
}

func TestDb_GetBudget(t *testing.T) {
	type args struct {
		ctx      context.Context
		budgetID uint64
	}
	tests := []struct {
		name                    string
		args                    args
		wantErr                 bool
		shouldCreateBankAccount bool
	}{
		{
			name: "[success] - delete budget for an existing milestone",
			args: args{
				ctx:      context.Background(),
				budgetID: uint64(helper.GenerateRandomId(10000, 300000)),
			},
			shouldCreateBankAccount: true,
			wantErr:                 false,
		},
		{
			name: "[failure] - delete budget for an invalid milestone",
			args: args{
				ctx:      context.Background(),
				budgetID: uint64(helper.GenerateRandomId(10000, 300000)),
			},
			wantErr:                 true,
			shouldCreateBankAccount: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldCreateBankAccount {
				sampleUserID := uint64(helper.GenerateRandomId(10000, 3000000))
				_, err := conn.CreateUserProfile(tt.args.ctx, &schema.UserProfile{
					UserId: sampleUserID,
				})

				assert.Nil(t, err)

				sampleBankAcct := helper.GenerateBankAccount()
				newAcct, err := conn.CreateBankAccount(tt.args.ctx, sampleUserID, sampleBankAcct)
				if err != nil {
					t.Errorf("Db.CreateBankAccount() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				// get a pocket from the acct
				// NOTE: this is guaranteed to work because we just created the bank account
				// which by default should have all the fields populated
				tt.args.budgetID = newAcct.Pockets[0].Goals[0].Milestones[0].Budget.Id
			}

			got, err := conn.GetBudget(tt.args.ctx, tt.args.budgetID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Db.GetBudget() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.NotNil(t, got)
				assert.Greater(t, got.Id, uint64(0))
				assert.Equal(t, got.Id, tt.args.budgetID)
			}
		})
	}
}

func TestDb_UpdateBudget(t *testing.T) {
	type args struct {
		ctx    context.Context
		budget *schema.Budget
	}
	tests := []struct {
		name                    string
		args                    args
		wantErr                 bool
		shouldCreateBankAccount bool
	}{
		{
			name: "[success] - delete budget for an existing milestone",
			args: args{
				ctx:    context.Background(),
				budget: helper.GenerateSingleBudget(),
			},
			shouldCreateBankAccount: true,
			wantErr:                 false,
		},
		{
			name: "[failure] - delete budget for an invalid milestone",
			args: args{
				ctx:    context.Background(),
				budget: helper.GenerateSingleBudget(),
			},
			wantErr:                 true,
			shouldCreateBankAccount: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			const budgetName = "new budget name"
			if tt.shouldCreateBankAccount {
				sampleUserID := uint64(helper.GenerateRandomId(10000, 3000000))
				_, err := conn.CreateUserProfile(tt.args.ctx, &schema.UserProfile{
					UserId: sampleUserID,
				})

				assert.Nil(t, err)

				sampleBankAcct := helper.GenerateBankAccount()
				newAcct, err := conn.CreateBankAccount(tt.args.ctx, sampleUserID, sampleBankAcct)
				if err != nil {
					t.Errorf("Db.CreateBankAccount() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				// get a pocket from the acct
				// NOTE: this is guaranteed to work because we just created the bank account
				// which by default should have all the fields populated
				tt.args.budget = newAcct.Pockets[0].Goals[0].Milestones[0].Budget
				// update the budget's name
				tt.args.budget.Name = budgetName
			}

			if err := conn.UpdateBudget(tt.args.ctx, tt.args.budget); (err != nil) != tt.wantErr {
				t.Errorf("Db.UpdateBudget() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				// check if the budget was updated
				updatedBudget, err := conn.GetBudget(tt.args.ctx, tt.args.budget.Id)
				assert.Nil(t, err)
				assert.Equal(t, updatedBudget.Name, budgetName)
			}
		})
	}
}
