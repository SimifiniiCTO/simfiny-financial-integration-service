package transformer

import (
	"reflect"
	"testing"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/plaid/plaid-go/v12/plaid"
)

func TestNewPlaidBankAccount(t *testing.T) {
	type args struct {
		userID      uint64
		bankAccount plaid.AccountBase
	}
	tests := []struct {
		name    string
		args    args
		want    *schema.BankAccount
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPlaidBankAccount(tt.args.userID, tt.args.bankAccount)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPlaidBankAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPlaidBankAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}
