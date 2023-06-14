package transformer

import (
	"reflect"
	"testing"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/plaid/plaid-go/v12/plaid"
)

func TestNewInvestmentAccount(t *testing.T) {
	type args struct {
		userID uint64
		input  *plaid.AccountBase
	}
	tests := []struct {
		name string
		args args
		want *schema.InvestmentAccount
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInvestmentAccount(tt.args.userID, tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInvestmentAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewInvestmentHolding(t *testing.T) {
	type args struct {
		input *plaid.Holding
	}
	tests := []struct {
		name string
		args args
		want *schema.InvesmentHolding
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInvestmentHolding(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInvestmentHolding() = %v, want %v", got, tt.want)
			}
		})
	}
}
