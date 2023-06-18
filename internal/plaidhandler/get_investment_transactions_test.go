package plaidhandler

import (
	"context"
	"testing"
	"time"

	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/helper"
	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/api/v1"
)

func TestPlaidWrapper_GetInvestmentTransactions(t *testing.T) {
	validLink := helper.GenerateLink(schema.LinkType_LINK_TYPE_PLAID)
	validLink.InvestmentAccounts = []*schema.InvestmentAccount{}
	validLink.Id = uint64(helper.GenerateRandomId(1000, 100000))

	type args struct {
		ctx                    context.Context
		accessToken            *string
		userId                 uint64
		link                   *schema.Link
		startDate              string
		endDate                string
		accountIds             []string
		numTransactionsToFetch int32
		offset                 int32
	}
	tests := []struct {
		name    string
		p       *PlaidWrapper
		args    args
		want    []*schema.InvestmentTransaction
		wantErr bool
	}{
		{
			name: "get investment transactions",
			p:    plaidTestClient,
			args: args{
				ctx:                    context.Background(),
				accessToken:            &testAccessToken,
				userId:                 uint64(helper.GenerateRandomId(1000, 100000)),
				startDate:              time.Now().AddDate(0, 0, -7).Format("2006-01-02"),
				endDate:                time.Now().Format("2006-01-02"),
				accountIds:             nil,
				numTransactionsToFetch: 0,
				offset:                 0,
				link:                   validLink,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.GetInvestmentTransactions(tt.args.ctx, tt.args.accessToken, &tt.args.userId, tt.args.link, tt.args.startDate, tt.args.endDate, tt.args.accountIds, tt.args.numTransactionsToFetch, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("PlaidWrapper.GetInvestmentTransactions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if len(got) == 0 {
					t.Errorf("PlaidWrapper.GetInvestmentTransactions() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
