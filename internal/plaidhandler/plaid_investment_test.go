package plaidhandler

import (
	"context"
	"reflect"
	"testing"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

func TestPlaidWrapper_GetInvestmentAccount(t *testing.T) {
	type args struct {
		ctx         context.Context
		userID      uint64
		accessToken string
	}
	tests := []struct {
		name    string
		p       *PlaidWrapper
		args    args
		want    []*schema.InvestmentAccount
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.GetInvestmentAccount(tt.args.ctx, tt.args.userID, tt.args.accessToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("PlaidWrapper.GetInvestmentAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlaidWrapper.GetInvestmentAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}
