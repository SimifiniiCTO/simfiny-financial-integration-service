package grpc

import (
	"context"
	"testing"
	"time"
)

func TestServer_DispatchPlaidSyncTask(t *testing.T) {
	type args struct {
		ctx         context.Context
		userId      uint64
		linkItemId  uint64
		accessToken string
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.DispatchPlaidSyncTask(tt.args.ctx, tt.args.userId, tt.args.linkItemId, tt.args.accessToken); (err != nil) != tt.wantErr {
				t.Errorf("Server.DispatchPlaidSyncTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServer_DispatchPullTransactionsTask(t *testing.T) {
	type args struct {
		ctx         context.Context
		userId      uint64
		linkId      uint64
		accessToken string
		startTime   time.Time
		endTime     time.Time
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.DispatchPullTransactionsTask(tt.args.ctx, tt.args.userId, tt.args.linkId, tt.args.accessToken, tt.args.startTime, tt.args.endTime); (err != nil) != tt.wantErr {
				t.Errorf("Server.DispatchPullTransactionsTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServer_DispatchPullUpdatedReCurringTransactionsTask(t *testing.T) {
	type args struct {
		ctx         context.Context
		userId      uint64
		linkId      uint64
		accessToken string
		accountIds  []string
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.DispatchPullUpdatedReCurringTransactionsTask(tt.args.ctx, tt.args.userId, tt.args.linkId, tt.args.accessToken, tt.args.accountIds); (err != nil) != tt.wantErr {
				t.Errorf("Server.DispatchPullUpdatedReCurringTransactionsTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServer_DispatchPullInvestmentTransactionsTask(t *testing.T) {
	type args struct {
		ctx         context.Context
		userId      uint64
		linkId      uint64
		accessToken string
		accountIds  []string
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.DispatchPullInvestmentTransactionsTask(tt.args.ctx, tt.args.userId, tt.args.linkId, tt.args.accessToken, tt.args.accountIds); (err != nil) != tt.wantErr {
				t.Errorf("Server.DispatchPullInvestmentTransactionsTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServer_DispatchPullInvestmentHoldingsTask(t *testing.T) {
	type args struct {
		ctx         context.Context
		userId      uint64
		linkId      uint64
		accessToken string
		accountIds  []string
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.DispatchPullInvestmentHoldingsTask(tt.args.ctx, tt.args.userId, tt.args.linkId, tt.args.accessToken, tt.args.accountIds); (err != nil) != tt.wantErr {
				t.Errorf("Server.DispatchPullInvestmentHoldingsTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
