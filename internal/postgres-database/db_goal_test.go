package database

import (
	"context"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/helper"
)

func TestDb_DeleteGoal(t *testing.T) {
	type args struct {
		ctx      context.Context
		pocketID uint64
		goal     *schema.SmartGoal
		goalID   uint64
	}
	tests := []struct {
		name                    string
		args                    args
		wantErr                 bool
		shouldCreateBankAccount bool
	}{
		{
			name: "delete goal",
			args: args{
				ctx:      context.Background(),
				goal:     helper.GenerateSingleGoal(3),
				pocketID: uint64(helper.GenerateRandomId(10000, 500000)),
				goalID:   uint64(helper.GenerateRandomId(10000, 500000)),
			},
			shouldCreateBankAccount: true,
			wantErr:                 false,
		},
		{
			name: "delete goal with invalid pocket id",
			args: args{
				ctx:      context.Background(),
				goal:     helper.GenerateSingleGoal(3),
				pocketID: uint64(helper.GenerateRandomId(10000, 500000)),
				goalID:   uint64(helper.GenerateRandomId(10000, 500000)),
			},
			shouldCreateBankAccount: false,
			wantErr:                 true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldCreateBankAccount {
				// create a bank account
				_, bankAcct := createUserProfileAndBankAccountForTest(t, tt.args.ctx, conn)
				// get the created pocket
				pocket := bankAcct.Pockets[0]
				tt.args.goal = pocket.Goals[0]
				tt.args.goalID = tt.args.goal.Id
			}

			if err := conn.DeleteGoal(tt.args.ctx, tt.args.goalID); (err != nil) != tt.wantErr {
				t.Errorf("Db.DeleteGoal() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				// check if the goal was deleted
				goal, err := conn.GetGoal(tt.args.ctx, tt.args.goalID)
				assert.Error(t, err)
				assert.Nil(t, goal)
			}
		})
	}
}

func TestDb_GetGoal(t *testing.T) {
	type args struct {
		ctx      context.Context
		pocketID uint64
		goal     *schema.SmartGoal
		goalID   uint64
	}
	tests := []struct {
		name                    string
		args                    args
		wantErr                 bool
		shouldCreateBankAccount bool
	}{
		{
			name: "get goal",
			args: args{
				ctx:      context.Background(),
				goal:     helper.GenerateSingleGoal(3),
				pocketID: uint64(helper.GenerateRandomId(10000, 500000)),
				goalID:   uint64(helper.GenerateRandomId(10000, 500000)),
			},
			shouldCreateBankAccount: true,
			wantErr:                 false,
		},
		{
			name: "get goal with invalid pocket id",
			args: args{
				ctx:      context.Background(),
				goal:     helper.GenerateSingleGoal(3),
				pocketID: uint64(helper.GenerateRandomId(10000, 500000)),
				goalID:   uint64(helper.GenerateRandomId(10000, 500000)),
			},
			shouldCreateBankAccount: false,
			wantErr:                 true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldCreateBankAccount {
				// create a bank account
				_, bankAcct := createUserProfileAndBankAccountForTest(t, tt.args.ctx, conn)
				// get the created pocket
				pocket := bankAcct.Pockets[0]
				tt.args.goal = pocket.Goals[0]
				tt.args.goalID = tt.args.goal.Id
			}
			got, err := conn.GetGoal(tt.args.ctx, tt.args.goalID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Db.GetGoal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.NotNil(t, got)
				assert.True(t, reflect.DeepEqual(got, tt.args.goal))
			}
		})
	}
}

func TestDb_GetGoalsByPocketID(t *testing.T) {
	type args struct {
		ctx      context.Context
		pocketID uint64
		goals    []*schema.SmartGoal
	}
	tests := []struct {
		name                    string
		args                    args
		wantErr                 bool
		shouldCreateBankAccount bool
	}{
		{
			name: "get goal",
			args: args{
				ctx:      context.Background(),
				goals:    helper.GenerateGoals(3),
				pocketID: uint64(helper.GenerateRandomId(10000, 500000)),
			},
			shouldCreateBankAccount: true,
			wantErr:                 false,
		},
		{
			name: "get goal with invalid pocket id",
			args: args{
				ctx:      context.Background(),
				goals:    helper.GenerateGoals(3),
				pocketID: uint64(helper.GenerateRandomId(10000, 500000)),
			},
			shouldCreateBankAccount: false,
			wantErr:                 true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldCreateBankAccount {
				// create a bank account
				_, bankAcct := createUserProfileAndBankAccountForTest(t, tt.args.ctx, conn)
				// get the created pocket
				pocket := bankAcct.Pockets[0]
				tt.args.goals = pocket.Goals
				tt.args.pocketID = pocket.Id
			}

			got, err := conn.GetGoalsByPocketID(tt.args.ctx, tt.args.pocketID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Db.GetGoalsByPocketID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.NotNil(t, got)
				assert.Equal(t, len(got), len(tt.args.goals))
			}
		})
	}
}

func TestDb_UpdateGoal(t *testing.T) {
	type args struct {
		ctx      context.Context
		pocketID uint64
		goal     *schema.SmartGoal
	}
	tests := []struct {
		name                    string
		args                    args
		wantErr                 bool
		shouldCreateBankAccount bool
	}{
		{
			name: "update goal",
			args: args{
				ctx:      context.Background(),
				goal:     helper.GenerateSingleGoal(3),
				pocketID: uint64(helper.GenerateRandomId(10000, 500000)),
			},
			shouldCreateBankAccount: true,
			wantErr:                 false,
		},
		{
			name: "update goal with invalid pocket id",
			args: args{
				ctx:      context.Background(),
				goal:     helper.GenerateSingleGoal(3),
				pocketID: uint64(helper.GenerateRandomId(10000, 500000)),
			},
			shouldCreateBankAccount: false,
			wantErr:                 true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			const updatedGoalName = "updated goal name"
			if tt.shouldCreateBankAccount {
				// create a bank account
				_, bankAcct := createUserProfileAndBankAccountForTest(t, tt.args.ctx, conn)
				// get the created pocket
				pocket := bankAcct.Pockets[0]
				tt.args.goal = pocket.Goals[0]
				tt.args.goal.Name = updatedGoalName
				tt.args.pocketID = pocket.Id
			}

			if err := conn.UpdateGoal(tt.args.ctx, tt.args.goal); (err != nil) != tt.wantErr {
				t.Errorf("Db.UpdateGoal() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				// get the updated goal
				updatedGoal, err := conn.GetGoal(tt.args.ctx, tt.args.goal.Id)
				assert.Nil(t, err)
				assert.Equal(t, updatedGoalName, updatedGoal.Name)
			}
		})
	}
}

func TestDb_CreateGoal(t *testing.T) {
	type args struct {
		ctx      context.Context
		pocketID uint64
		goal     *schema.SmartGoal
	}
	tests := []struct {
		name                    string
		args                    args
		wantErr                 bool
		shouldCreateBankAccount bool
	}{
		{
			name: "create goal",
			args: args{
				ctx:      context.Background(),
				goal:     helper.GenerateSingleGoal(3),
				pocketID: uint64(helper.GenerateRandomId(10000, 500000)),
			},
			shouldCreateBankAccount: true,
			wantErr:                 false,
		},
		{
			name: "create goal with invalid pocket id",
			args: args{
				ctx:      context.Background(),
				goal:     helper.GenerateSingleGoal(3),
				pocketID: uint64(helper.GenerateRandomId(10000, 500000)),
			},
			shouldCreateBankAccount: false,
			wantErr:                 true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldCreateBankAccount {
				// create a bank account
				_, bankAcct := createUserProfileAndBankAccountForTest(t, tt.args.ctx, conn)
				// get the created pocket
				pocket := bankAcct.Pockets[0]
				tt.args.pocketID = pocket.Id
			}

			got, err := conn.CreateGoal(tt.args.ctx, tt.args.pocketID, tt.args.goal)
			if (err != nil) != tt.wantErr {
				t.Errorf("Db.CreateGoal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.NotNil(t, got)
			}
		})
	}
}
