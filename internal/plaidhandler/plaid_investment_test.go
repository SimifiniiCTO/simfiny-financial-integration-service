package plaidhandler

import (
	"context"
	"testing"

	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/helper"
	"github.com/stretchr/testify/assert"
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
		wantErr bool
	}{
		{
			name: "get investment account",
			p:    plaidTestClient,
			args: args{
				ctx:         context.Background(),
				userID:      uint64(helper.GenerateRandomId(1000, 100000)),
				accessToken: testAccessToken,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.GetInvestmentAccount(tt.args.ctx, tt.args.userID, tt.args.accessToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("PlaidWrapper.GetInvestmentAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.NotNil(t, got)
			}
		})
	}
}
