package plaidhandler

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPlaidWrapper_GetAllTransactions(t *testing.T) {
	type args struct {
		ctx         context.Context
		accessToken string
		start       time.Time
		end         time.Time
		accountIds  []string
	}
	tests := []struct {
		name    string
		p       *PlaidWrapper
		args    args
		wantErr bool
	}{
		{
			name: "get past 7 days of transactions",
			p:    plaidTestClient,
			args: args{
				ctx:         context.Background(),
				accessToken: testAccessToken,
				start:       time.Now().AddDate(0, 0, -7),
				end:         time.Now(),
				accountIds:  []string{},
			},
			wantErr: false,
		},
		{
			name: "get past 31 days of transactions",
			p:    plaidTestClient,
			args: args{
				ctx:         context.Background(),
				accessToken: testAccessToken,
				start:       time.Now().AddDate(0, 0, -31),
				end:         time.Now(),
				accountIds:  []string{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.GetAllTransactions(tt.args.ctx, tt.args.accessToken, tt.args.start, tt.args.end, tt.args.accountIds)
			if (err != nil) != tt.wantErr {
				t.Errorf("PlaidWrapper.GetAllTransactions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				assert.NotNil(t, got)
			}
		})
	}
}
