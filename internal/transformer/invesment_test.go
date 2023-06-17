package transformer

import (
	"testing"

	"github.com/plaid/plaid-go/v12/plaid"
	"github.com/stretchr/testify/assert"
)

func TestNewInvestmentAccount(t *testing.T) {
	type args struct {
		userID uint64
		input  *plaid.AccountBase
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid investment account",
			args: args{
				userID: 1,
				input:  generateSinglePlaidAccountBase(plaid.ACCOUNTTYPE_INVESTMENT),
			},
			wantErr: false,
		},
		{
			name: "invalid investment account",
			args: args{
				userID: 1,
				input:  generateSinglePlaidAccountBase(plaid.ACCOUNTTYPE_BROKERAGE),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewInvestmentAccount(tt.args.userID, tt.args.input)
			if err != nil && !tt.wantErr {
				t.Errorf("NewInvestmentAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.NotNil(t, got)
			}
		})
	}
}

func TestNewInvestmentHolding(t *testing.T) {
	type args struct {
		input *plaid.Holding
	}
	tests := []struct {
		name      string
		args      args
		wantError bool
	}{
		{
			name: "valid investment holding",
			args: args{
				input: generateInvestmentHolding(),
			},
			wantError: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewInvestmentHolding(tt.args.input)
			if got != nil && tt.wantError {
				t.Errorf("NewInvestmentHolding() = %v, want %v", got, tt.wantError)
			}
		})
	}
}
