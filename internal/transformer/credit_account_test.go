package transformer

import (
	"testing"

	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/helper"
	"github.com/plaid/plaid-go/v12/plaid"
	"github.com/stretchr/testify/assert"
)

func TestNewCreditAccount(t *testing.T) {
	type args struct {
		userID uint64
		input  *plaid.CreditCardLiability
		acct   *plaid.AccountBase
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid credit account",
			args: args{
				userID: uint64(helper.GenerateRandomId(1000, 1000000)),
				input:  generateSingleCreditCardLiability(),
				acct:   generateSinglePlaidAccountBase(plaid.ACCOUNTTYPE_CREDIT),
			},
			wantErr: false,
		},
		{
			name: "invalid credit account",
			args: args{
				userID: uint64(helper.GenerateRandomId(1000, 1000000)),
				input:  generateSingleCreditCardLiability(),
				acct:   generateSinglePlaidAccountBase(plaid.ACCOUNTTYPE_BROKERAGE),
			},
			wantErr: true,
		},
		{
			name: "invalid credit account",
			args: args{
				userID: uint64(helper.GenerateRandomId(1000, 1000000)),
				input:  generateSingleCreditCardLiability(),
				acct:   generateSinglePlaidAccountBase(plaid.ACCOUNTTYPE_DEPOSITORY),
			},
			wantErr: true,
		},
		{
			name: "invalid credit account",
			args: args{
				userID: uint64(helper.GenerateRandomId(1000, 1000000)),
				input:  generateSingleCreditCardLiability(),
				acct:   generateSinglePlaidAccountBase(plaid.ACCOUNTTYPE_INVESTMENT),
			},
			wantErr: true,
		},
		{
			name: "invalid credit account",
			args: args{
				userID: uint64(helper.GenerateRandomId(1000, 1000000)),
				input:  generateSingleCreditCardLiability(),
				acct:   generateSinglePlaidAccountBase(plaid.ACCOUNTTYPE_LOAN),
			},
			wantErr: true,
		},
		{
			name: "invalid credit account",
			args: args{
				userID: uint64(helper.GenerateRandomId(1000, 1000000)),
				input:  generateSingleCreditCardLiability(),
				acct:   nil,
			},
			wantErr: true,
		},
		{
			name: "invalid credit account",
			args: args{
				userID: uint64(helper.GenerateRandomId(1000, 1000000)),
				input:  nil,
				acct:   nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCreditAccount(tt.args.userID, tt.args.input, tt.args.acct)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCreditAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.NotNil(t, got)
			}
		})
	}
}

func TestNewAPR(t *testing.T) {
	type args struct {
		plaidAPR []plaid.APR
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid apr",
			args: args{
				plaidAPR: generateMultiplePlaidAPR(10),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAPR(&tt.args.plaidAPR)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAPR() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.NotNil(t, got)
			}

		})
	}
}
