package plaidhandler

import (
	"context"
	"reflect"
	"testing"
	"time"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

func TestPlaidWrapper_GetAllTransactions(t *testing.T) {
	type args struct {
		ctx         context.Context
		accessToken string
		start       time.Time
		end         time.Time
		accountIds  []string
	}
	tests := []struct {
		name    string
		p       *PlaidWrapper
		args    args
		want    []*schema.Transaction
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.GetAllTransactions(tt.args.ctx, tt.args.accessToken, tt.args.start, tt.args.end, tt.args.accountIds)
			if (err != nil) != tt.wantErr {
				t.Errorf("PlaidWrapper.GetAllTransactions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlaidWrapper.GetAllTransactions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlaidWrapper_GetTransactions(t *testing.T) {
	type args struct {
		ctx            context.Context
		accessToken    string
		start          time.Time
		end            time.Time
		count          int32
		offset         int32
		bankAccountIds []string
	}
	tests := []struct {
		name    string
		p       *PlaidWrapper
		args    args
		want    []*schema.Transaction
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.GetTransactions(tt.args.ctx, tt.args.accessToken, tt.args.start, tt.args.end, tt.args.count, tt.args.offset, tt.args.bankAccountIds)
			if (err != nil) != tt.wantErr {
				t.Errorf("PlaidWrapper.GetTransactions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlaidWrapper.GetTransactions() = %v, want %v", got, tt.want)
			}
		})
	}
}
