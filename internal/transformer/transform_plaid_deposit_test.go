package transformer

import (
	"reflect"
	"testing"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/plaid/plaid-go/v12/plaid"
)

func Test_transformPlaidAccountToDepositAccount(t *testing.T) {
	type args struct {
		accounts *[]plaid.AccountBase
	}
	tests := []struct {
		name string
		args args
		want []*schema.BankAccount
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := transformPlaidAccountToDepositAccount(tt.args.accounts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transformPlaidAccountToDepositAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}
