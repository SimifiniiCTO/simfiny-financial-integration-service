package transformer

import (
	"reflect"
	"testing"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/plaid/plaid-go/v12/plaid"
)

func Test_transformPlaidInvestmentObject(t *testing.T) {
	type args struct {
		securities      *[]plaid.Security
		holdings        *[]plaid.Holding
		acctIDToTypeMap map[string]*accountMetadata
	}
	tests := []struct {
		name    string
		args    args
		want    []*schema.InvestmentAccount
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := transformPlaidInvestmentObject(tt.args.securities, tt.args.holdings, tt.args.acctIDToTypeMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("transformPlaidInvestmentObject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transformPlaidInvestmentObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transformAndMapSecuritiesToAcctType(t *testing.T) {
	type args struct {
		securities *[]plaid.Security
	}
	tests := []struct {
		name string
		args args
		want map[string][]*schema.InvestmentSecurity
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := transformAndMapSecuritiesToAcctType(tt.args.securities); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transformAndMapSecuritiesToAcctType() = %v, want %v", got, tt.want)
			}
		})
	}
}
