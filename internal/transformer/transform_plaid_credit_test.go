package transformer

import (
	"reflect"
	"testing"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/plaid/plaid-go/v12/plaid"
)

func TestTransformPlaidCredit(t *testing.T) {
	type args struct {
		creditCardliabilities *[]plaid.CreditCardLiability
		acctIDToTypeMap       map[string]*accountMetadata
	}
	tests := []struct {
		name    string
		args    args
		want    []*schema.CreditAccount
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TransformPlaidCredit(tt.args.creditCardliabilities, tt.args.acctIDToTypeMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("TransformPlaidCredit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransformPlaidCredit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transformPlaidApr(t *testing.T) {
	type args struct {
		plaidAPR []plaid.APR
	}
	tests := []struct {
		name string
		args args
		want []*schema.Apr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := transformPlaidApr(tt.args.plaidAPR); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transformPlaidApr() = %v, want %v", got, tt.want)
			}
		})
	}
}
