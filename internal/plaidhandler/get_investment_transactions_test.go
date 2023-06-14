package plaidhandler

import (
	"context"
	"reflect"
	"testing"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/plaid/plaid-go/v12/plaid"
)

func TestPlaidWrapper_GetInvestmentTransactions(t *testing.T) {
	type args struct {
		ctx                    context.Context
		accessToken            *string
		userId                 *uint64
		link                   *schema.Link
		startDate              string
		endDate                string
		accountIds             []string
		numTransactionsToFetch int32
		offset                 int32
	}
	tests := []struct {
		name    string
		p       *PlaidWrapper
		args    args
		want    []*schema.InvestmentTransaction
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.GetInvestmentTransactions(tt.args.ctx, tt.args.accessToken, tt.args.userId, tt.args.link, tt.args.startDate, tt.args.endDate, tt.args.accountIds, tt.args.numTransactionsToFetch, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("PlaidWrapper.GetInvestmentTransactions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlaidWrapper.GetInvestmentTransactions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlaidWrapper_getInvestmentTransactions(t *testing.T) {
	type args struct {
		ctx    context.Context
		req    *plaid.InvestmentsTransactionsGetRequest
		userId *uint64
		link   *schema.Link
	}
	tests := []struct {
		name    string
		p       *PlaidWrapper
		args    args
		want    []*schema.InvestmentTransaction
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.getInvestmentTransactions(tt.args.ctx, tt.args.req, tt.args.userId, tt.args.link)
			if (err != nil) != tt.wantErr {
				t.Errorf("PlaidWrapper.getInvestmentTransactions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlaidWrapper.getInvestmentTransactions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toInvestmentTransactions(t *testing.T) {
	type args struct {
		tx     []plaid.InvestmentTransaction
		linkId uint64
		userId uint64
	}
	tests := []struct {
		name string
		args args
		want []*schema.InvestmentTransaction
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toInvestmentTransactions(tt.args.tx, tt.args.linkId, tt.args.userId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toInvestmentTransactions() = %v, want %v", got, tt.want)
			}
		})
	}
}
