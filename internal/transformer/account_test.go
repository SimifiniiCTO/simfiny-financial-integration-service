package transformer

import (
	"testing"

	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/helper"
	"github.com/plaid/plaid-go/v14/plaid"
	"github.com/stretchr/testify/assert"
)

func TestNewPlaidBankAccount(t *testing.T) {
	type args struct {
		userID      uint64
		bankAccount *plaid.AccountBase
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid bank account",
			args: args{
				userID:      uint64(helper.GenerateRandomId(1000, 1000000)),
				bankAccount: generateSinglePlaidAccountBase(plaid.ACCOUNTTYPE_DEPOSITORY),
			},
			wantErr: false,
		},
		{
			name: "invalid bank account",
			args: args{
				userID:      uint64(helper.GenerateRandomId(1000, 1000000)),
				bankAccount: generateSinglePlaidAccountBase(plaid.ACCOUNTTYPE_CREDIT),
			},
			wantErr: true,
		},
		{
			name: "invalid bank account",
			args: args{
				userID:      uint64(helper.GenerateRandomId(1000, 1000000)),
				bankAccount: generateSinglePlaidAccountBase(plaid.ACCOUNTTYPE_BROKERAGE),
			},
			wantErr: true,
		},
		{
			name: "invalid bank account",
			args: args{
				userID:      uint64(helper.GenerateRandomId(1000, 1000000)),
				bankAccount: generateSinglePlaidAccountBase(plaid.ACCOUNTTYPE_LOAN),
			},
			wantErr: true,
		},
		{
			name: "invalid bank account",
			args: args{
				userID:      uint64(helper.GenerateRandomId(1000, 1000000)),
				bankAccount: generateSinglePlaidAccountBase(plaid.ACCOUNTTYPE_OTHER),
			},
			wantErr: true,
		},
		{
			name: "invalid bank account",
			args: args{
				userID:      uint64(helper.GenerateRandomId(1000, 1000000)),
				bankAccount: generateSinglePlaidAccountBase(plaid.ACCOUNTTYPE_INVESTMENT),
			},
			wantErr: true,
		},
		{
			name: "invalid bank account",
			args: args{
				userID:      uint64(helper.GenerateRandomId(1000, 1000000)),
				bankAccount: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPlaidBankAccount(tt.args.userID, tt.args.bankAccount)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPlaidBankAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.NotNil(t, got)
				assert.True(t, got.UserId == tt.args.userID)
			}
		})
	}
}
