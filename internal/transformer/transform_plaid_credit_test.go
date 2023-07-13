package transformer

import (
	"testing"

	"github.com/plaid/plaid-go/v14/plaid"
	"github.com/stretchr/testify/assert"
)

func TestTransformPlaidCredit(t *testing.T) {
	creditCardLiabilities, acctMeta := generateManyPlaidCreditLiabilityAndAccountMetadata(10)
	type args struct {
		creditCardliabilities *[]plaid.CreditCardLiability
		acctIDToTypeMap       map[string]*accountMetadata
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid credit card liability",
			args: args{
				creditCardliabilities: &creditCardLiabilities,
				acctIDToTypeMap:       acctMeta,
			},
			wantErr: false,
		},
		{
			name: "invalid credit card liability",
			args: args{
				creditCardliabilities: nil,
				acctIDToTypeMap:       acctMeta,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TransformPlaidCredit(tt.args.creditCardliabilities, tt.args.acctIDToTypeMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("TransformPlaidCredit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.NotNil(t, got)
			}
		})
	}
}
