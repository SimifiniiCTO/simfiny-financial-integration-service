package transformer

import (
	"reflect"
	"testing"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/plaid/plaid-go/v12/plaid"
)

func TestTransformPlaidInvestmentSecurity(t *testing.T) {
	type args struct {
		securities []plaid.Security
	}
	tests := []struct {
		name    string
		args    args
		want    []*schema.InvestmentSecurity
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TransformPlaidInvestmentSecurity(tt.args.securities)
			if (err != nil) != tt.wantErr {
				t.Errorf("TransformPlaidInvestmentSecurity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransformPlaidInvestmentSecurity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransformSinglePlaidInvestmentSecurity(t *testing.T) {
	type args struct {
		security plaid.Security
	}
	tests := []struct {
		name string
		args args
		want *schema.InvestmentSecurity
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TransformSinglePlaidInvestmentSecurity(tt.args.security); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransformSinglePlaidInvestmentSecurity() = %v, want %v", got, tt.want)
			}
		})
	}
}
