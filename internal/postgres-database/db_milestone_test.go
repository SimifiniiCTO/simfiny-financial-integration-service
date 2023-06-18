package database

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/helper"
	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/api/v1"
)

func TestDb_UpdateMilestone(t *testing.T) {
	type args struct {
		ctx         context.Context
		goalID      uint64
		milestone   *schema.Milestone
		milestoneID uint64
	}
	tests := []struct {
		name                    string
		args                    args
		wantErr                 bool
		shouldCreateBankAccount bool
	}{
		{
			name: "Update milestone",
			args: args{
				ctx:         context.Background(),
				goalID:      uint64(helper.GenerateRandomId(10000, 50000)),
				milestone:   helper.GenerateSingleMilestone(3),
				milestoneID: uint64(helper.GenerateRandomId(10000, 50000)),
			},
			wantErr:                 false,
			shouldCreateBankAccount: true,
		},
		{
			name: "Update milestone - GoalID not found",
			args: args{
				ctx:         context.Background(),
				goalID:      uint64(helper.GenerateRandomId(10000, 50000)),
				milestone:   helper.GenerateSingleMilestone(3),
				milestoneID: uint64(helper.GenerateRandomId(10000, 50000)),
			},
			wantErr:                 true,
			shouldCreateBankAccount: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			const milestoneName = "Updated milestone name"
			if tt.shouldCreateBankAccount {
				_, bankAcct := createUserProfileAndBankAccountForTest(t, tt.args.ctx, conn)
				// get the goal
				goal := bankAcct.Pockets[0].Goals[0]
				tt.args.goalID = goal.Id
				tt.args.milestone = goal.Milestones[0]
				tt.args.milestoneID = tt.args.milestone.Id

				// update the milestone name
				tt.args.milestone.Name = milestoneName
			}

			if err := conn.UpdateMilestone(tt.args.ctx, tt.args.milestone); (err != nil) != tt.wantErr {
				t.Errorf("Db.UpdateMilestone() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				// get the milestone
				milestone, err := conn.GetMilestone(tt.args.ctx, tt.args.milestoneID)
				assert.NoError(t, err)
				assert.Equal(t, milestoneName, milestone.Name)
			}
		})
	}
}

func TestDb_GetMilestone(t *testing.T) {
	type args struct {
		ctx         context.Context
		goalID      uint64
		milestone   *schema.Milestone
		milestoneID uint64
	}
	tests := []struct {
		name                    string
		args                    args
		wantErr                 bool
		shouldCreateBankAccount bool
	}{
		{
			name: "Get milestone",
			args: args{
				ctx:       context.Background(),
				goalID:    uint64(helper.GenerateRandomId(10000, 50000)),
				milestone: helper.GenerateSingleMilestone(3),
			},
			wantErr:                 false,
			shouldCreateBankAccount: true,
		},
		{
			name: "Get milestone - GoalID not found",
			args: args{
				ctx:       context.Background(),
				goalID:    uint64(helper.GenerateRandomId(10000, 50000)),
				milestone: helper.GenerateSingleMilestone(3),
			},
			wantErr:                 true,
			shouldCreateBankAccount: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldCreateBankAccount {
				_, bankAcct := createUserProfileAndBankAccountForTest(t, tt.args.ctx, conn)
				// get the goal
				goal := bankAcct.Pockets[0].Goals[0]
				tt.args.goalID = goal.Id
				tt.args.milestone = goal.Milestones[0]
				tt.args.milestoneID = tt.args.milestone.Id
			}
			got, err := conn.GetMilestone(tt.args.ctx, tt.args.milestoneID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Db.GetMilestone() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.Equal(t, tt.args.milestone, got)
			}
		})
	}
}

func TestDb_GetMilestonesByGoalID(t *testing.T) {
	type args struct {
		ctx        context.Context
		goalID     uint64
		milestones []*schema.Milestone
	}
	tests := []struct {
		name                    string
		args                    args
		wantErr                 bool
		shouldCreateBankAccount bool
	}{
		{
			name: "Get milestones",
			args: args{
				ctx:        context.Background(),
				goalID:     uint64(helper.GenerateRandomId(10000, 50000)),
				milestones: helper.GenerateMilestones(3),
			},
			wantErr:                 false,
			shouldCreateBankAccount: true,
		},
		{
			name: "Get milestones - GoalID not found",
			args: args{
				ctx:        context.Background(),
				goalID:     uint64(helper.GenerateRandomId(10000, 50000)),
				milestones: helper.GenerateMilestones(3),
			},
			wantErr:                 true,
			shouldCreateBankAccount: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldCreateBankAccount {
				_, bankAcct := createUserProfileAndBankAccountForTest(t, tt.args.ctx, conn)
				// get the goal
				goal := bankAcct.Pockets[0].Goals[0]
				tt.args.goalID = goal.Id
				tt.args.milestones = goal.Milestones
			}

			got, err := conn.GetMilestonesByGoalID(tt.args.ctx, tt.args.goalID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Db.GetMilestonesByGoalID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.Equal(t, len(tt.args.milestones), len(got))
			}
		})
	}
}

func TestDb_DeleteMilestone(t *testing.T) {
	type args struct {
		ctx         context.Context
		goalID      uint64
		milestone   *schema.Milestone
		milestoneID uint64
	}
	tests := []struct {
		name                    string
		args                    args
		wantErr                 bool
		shouldCreateBankAccount bool
	}{
		{
			name: "Delete milestone",
			args: args{
				ctx:       context.Background(),
				goalID:    uint64(helper.GenerateRandomId(10000, 50000)),
				milestone: helper.GenerateSingleMilestone(3),
			},
			wantErr:                 false,
			shouldCreateBankAccount: true,
		},
		{
			name: "Delete milestone - GoalID not found",
			args: args{
				ctx:       context.Background(),
				goalID:    uint64(helper.GenerateRandomId(10000, 50000)),
				milestone: helper.GenerateSingleMilestone(3),
			},
			wantErr:                 true,
			shouldCreateBankAccount: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldCreateBankAccount {
				_, bankAcct := createUserProfileAndBankAccountForTest(t, tt.args.ctx, conn)
				// get the goal
				goal := bankAcct.Pockets[0].Goals[0]
				tt.args.goalID = goal.Id
				tt.args.milestone = goal.Milestones[0]
				tt.args.milestoneID = tt.args.milestone.Id
			}

			if err := conn.DeleteMilestone(tt.args.ctx, tt.args.milestoneID); (err != nil) != tt.wantErr {
				t.Errorf("Db.DeleteMilestone() error = %v, wantErr %v", err, tt.wantErr)
			}

			// check if the milestone is deleted
			if !tt.wantErr {
				_, err := conn.GetMilestone(tt.args.ctx, tt.args.milestoneID)
				assert.Error(t, err)
			}
		})
	}
}

func TestDb_CreateMilestone(t *testing.T) {
	type args struct {
		ctx       context.Context
		goalID    uint64
		milestone *schema.Milestone
	}
	tests := []struct {
		name                    string
		args                    args
		wantErr                 bool
		shouldCreateBankAccount bool
	}{
		{
			name: "CreateMilestone",
			args: args{
				ctx:       context.Background(),
				goalID:    uint64(helper.GenerateRandomId(10000, 50000)),
				milestone: helper.GenerateSingleMilestone(3),
			},
			wantErr:                 false,
			shouldCreateBankAccount: true,
		},
		{
			name: "CreateMilestone - GoalID not found",
			args: args{
				ctx:       context.Background(),
				goalID:    uint64(helper.GenerateRandomId(10000, 50000)),
				milestone: helper.GenerateSingleMilestone(3),
			},
			wantErr:                 true,
			shouldCreateBankAccount: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldCreateBankAccount {
				_, bankAcct := createUserProfileAndBankAccountForTest(t, tt.args.ctx, conn)
				// get the goal
				goal := bankAcct.Pockets[0].Goals[0]
				tt.args.goalID = goal.Id
				tt.args.milestone = goal.Milestones[0]
			}

			got, err := conn.CreateMilestone(tt.args.ctx, tt.args.goalID, tt.args.milestone)
			if (err != nil) != tt.wantErr {
				t.Errorf("Db.CreateMilestone() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.NotNil(t, got)
				assert.Equal(t, tt.args.milestone.Name, got.Name)
			}
		})
	}
}
