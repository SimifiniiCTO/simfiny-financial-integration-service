package transformer

import (
	"reflect"
	"testing"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/plaid/plaid-go/v12/plaid"
)

func TestTransformMortgageAccount(t *testing.T) {
	type args struct {
		mortgageLoan    *[]plaid.MortgageLiability
		acctIDToTypeMap map[string]*accountMetadata
	}
	tests := []struct {
		name    string
		args    args
		want    []*schema.MortgageAccount
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TransformMortgageAccount(tt.args.mortgageLoan, tt.args.acctIDToTypeMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("TransformMortgageAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransformMortgageAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}
