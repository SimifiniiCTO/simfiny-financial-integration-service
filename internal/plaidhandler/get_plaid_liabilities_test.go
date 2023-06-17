package plaidhandler

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlaidWrapper_GetPlaidLiabilityAccounts(t *testing.T) {

	type args struct {
		ctx         context.Context
		accessToken *string
	}
	tests := []struct {
		name    string
		p       *PlaidWrapper
		args    args
		wantErr bool
	}{
		{
			name: "get plaid liability accounts",
			p:    plaidTestClient,
			args: args{
				ctx:         context.Background(),
				accessToken: &testAccessToken,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.GetPlaidLiabilityAccounts(tt.args.ctx, tt.args.accessToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("PlaidWrapper.GetPlaidLiabilityAccounts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.NotNil(t, got)
			}
		})
	}
}
