package taskhandler

import (
	"context"
	"reflect"
	"testing"

	"github.com/SimifiniiCTO/asynq"
)

func TestNewPullInvestmentHoldingsTask(t *testing.T) {
	type args struct {
		userId      uint64
		linkId      uint64
		accessToken string
		accountIds  []string
	}
	tests := []struct {
		name    string
		args    args
		want    *asynq.Task
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPullInvestmentHoldingsTask(tt.args.userId, tt.args.linkId, tt.args.accessToken, tt.args.accountIds)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPullInvestmentHoldingsTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPullInvestmentHoldingsTask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskHandler_RunPullInvestmentHoldingsTask(t *testing.T) {
	type args struct {
		ctx context.Context
		t   *asynq.Task
	}
	tests := []struct {
		name    string
		th      *TaskHandler
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.th.RunPullInvestmentHoldingsTask(tt.args.ctx, tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("TaskHandler.RunPullInvestmentHoldingsTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_isInAccountIdSlice(t *testing.T) {
	type args struct {
		accountId  string
		accountIds []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isInAccountIdSlice(tt.args.accountId, tt.args.accountIds); got != tt.want {
				t.Errorf("isInAccountIdSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
