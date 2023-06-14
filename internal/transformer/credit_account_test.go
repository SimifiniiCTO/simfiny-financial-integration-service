package transformer

import (
	"reflect"
	"testing"

	"github.com/SimifiniiCTO/simfiny-financial-integration-service/proto"
	"github.com/plaid/plaid-go/v12/plaid"
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
		want    *proto.CreditAccount
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCreditAccount(tt.args.userID, tt.args.input, tt.args.acct)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCreditAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCreditAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAPR(t *testing.T) {
	type args struct {
		plaidAPR *[]plaid.APR
	}
	tests := []struct {
		name    string
		args    args
		want    []*proto.APR
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAPR(tt.args.plaidAPR)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAPR() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAPR() = %v, want %v", got, tt.want)
			}
		})
	}
}
