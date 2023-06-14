package plaidhandler

import (
	"context"
	"reflect"
	"testing"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/plaid/plaid-go/v12/plaid"
)

func TestPlaidWrapper_GetRecurringTransactionsForAccounts(t *testing.T) {
	type args struct {
		ctx             context.Context
		accessToken     *string
		userId          *uint64
		linkId          *uint64
		plaidAccountIds []string
	}
	tests := []struct {
		name    string
		p       *PlaidWrapper
		args    args
		want    []*schema.ReOccuringTransaction
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.GetRecurringTransactionsForAccounts(tt.args.ctx, tt.args.accessToken, tt.args.userId, tt.args.linkId, tt.args.plaidAccountIds)
			if (err != nil) != tt.wantErr {
				t.Errorf("PlaidWrapper.GetRecurringTransactionsForAccounts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlaidWrapper.GetRecurringTransactionsForAccounts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlaidWrapper_GetRecurringTransactions(t *testing.T) {
	type args struct {
		ctx         context.Context
		accessToken *string
		userId      *uint64
		linkId      *uint64
	}
	tests := []struct {
		name    string
		p       *PlaidWrapper
		args    args
		want    []*schema.ReOccuringTransaction
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.GetRecurringTransactions(tt.args.ctx, tt.args.accessToken, tt.args.userId, tt.args.linkId)
			if (err != nil) != tt.wantErr {
				t.Errorf("PlaidWrapper.GetRecurringTransactions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlaidWrapper.GetRecurringTransactions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlaidWrapper_getRecurringTransactions(t *testing.T) {
	type args struct {
		ctx    context.Context
		req    *plaid.TransactionsRecurringGetRequest
		userId *uint64
		linkId *uint64
	}
	tests := []struct {
		name    string
		p       *PlaidWrapper
		args    args
		want    []*schema.ReOccuringTransaction
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.getRecurringTransactions(tt.args.ctx, tt.args.req, tt.args.userId, tt.args.linkId)
			if (err != nil) != tt.wantErr {
				t.Errorf("PlaidWrapper.getRecurringTransactions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlaidWrapper.getRecurringTransactions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionStreamToRecurringTransactions(t *testing.T) {
	type args struct {
		userId  *uint64
		linkId  *uint64
		streams []plaid.TransactionStream
		flow    schema.ReCurringFlow
	}
	tests := []struct {
		name string
		args args
		want []*schema.ReOccuringTransaction
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := transactionStreamToRecurringTransactions(tt.args.userId, tt.args.linkId, tt.args.streams, tt.args.flow); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transactionStreamToRecurringTransactions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getFrequency(t *testing.T) {
	type args struct {
		freq plaid.RecurringTransactionFrequency
	}
	tests := []struct {
		name string
		args args
		want schema.ReOccuringTransactionsFrequency
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFrequency(tt.args.freq); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFrequency() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getStatus(t *testing.T) {
	type args struct {
		status plaid.TransactionStreamStatus
	}
	tests := []struct {
		name string
		args args
		want schema.ReOccuringTransactionsStatus
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getStatus(tt.args.status); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}
