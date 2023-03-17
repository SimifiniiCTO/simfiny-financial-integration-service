package database

import (
	"context"
	"reflect"
	"testing"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

func TestDb_UpdateMilestone(t *testing.T) {
	type args struct {
		ctx       context.Context
		milestone *schema.Milestone
	}
	tests := []struct {
		name    string
		d       *Db
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.UpdateMilestone(tt.args.ctx, tt.args.milestone); (err != nil) != tt.wantErr {
				t.Errorf("Db.UpdateMilestone() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDb_GetMilestone(t *testing.T) {
	type args struct {
		ctx         context.Context
		milestoneID uint64
	}
	tests := []struct {
		name    string
		d       *Db
		args    args
		want    *schema.Milestone
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.d.GetMilestone(tt.args.ctx, tt.args.milestoneID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Db.GetMilestone() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Db.GetMilestone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDb_GetMilestonesByGoalID(t *testing.T) {
	type args struct {
		ctx    context.Context
		goalID uint64
	}
	tests := []struct {
		name    string
		d       *Db
		args    args
		want    []*schema.Milestone
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.d.GetMilestonesByGoalID(tt.args.ctx, tt.args.goalID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Db.GetMilestonesByGoalID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Db.GetMilestonesByGoalID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDb_DeleteMilestone(t *testing.T) {
	type args struct {
		ctx         context.Context
		milestoneID uint64
	}
	tests := []struct {
		name    string
		d       *Db
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.DeleteMilestone(tt.args.ctx, tt.args.milestoneID); (err != nil) != tt.wantErr {
				t.Errorf("Db.DeleteMilestone() error = %v, wantErr %v", err, tt.wantErr)
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
		name    string
		d       *Db
		args    args
		want    *schema.Milestone
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.d.CreateMilestone(tt.args.ctx, tt.args.goalID, tt.args.milestone)
			if (err != nil) != tt.wantErr {
				t.Errorf("Db.CreateMilestone() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Db.CreateMilestone() = %v, want %v", got, tt.want)
			}
		})
	}
}
