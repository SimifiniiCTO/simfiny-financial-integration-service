package transformer

import (
	"testing"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/plaid/plaid-go/v12/plaid"
	"github.com/stretchr/testify/assert"
)

func TestTransformMortgageAccount(t *testing.T) {
	mortgageLiabilities, _ := generateManyPlaidMortgageLiabilityAndAccountMetadata(10)

	type args struct {
		mortgageLoan *[]plaid.MortgageLiability
	}
	tests := []struct {
		name    string
		args    args
		want    []*schema.MortgageAccount
		wantErr bool
	}{
		{
			name: "valid mortgage loan liability",
			args: args{
				mortgageLoan: &mortgageLiabilities,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TransformMortgageAccount(tt.args.mortgageLoan)
			if (err != nil) != tt.wantErr {
				t.Errorf("TransformMortgageAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.NotNil(t, got)
			}

		})
	}
}
