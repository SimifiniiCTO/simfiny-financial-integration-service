package plaidhandler

import (
	"context"
	"reflect"
	"testing"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

func TestPlaidWrapper_GetAccounts(t *testing.T) {
	type args struct {
		ctx         context.Context
		accessToken string
		userId      uint64
	}
	tests := []struct {
		name    string
		p       *PlaidWrapper
		args    args
		want    []*schema.BankAccount
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.GetAccounts(tt.args.ctx, tt.args.accessToken, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("PlaidWrapper.GetAccounts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlaidWrapper.GetAccounts() = %v, want %v", got, tt.want)
			}
		})
	}
}
