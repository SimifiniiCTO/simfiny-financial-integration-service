package plaidhandler

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/plaid/plaid-go/v12/plaid"
	"github.com/stretchr/testify/assert"
)

type plaidInvestmentScenarios struct {
	scenarioName     string
	shouldErrorOccur bool
	accessToken      *string
	expectedError    error
}

// getPlaidInvesmentScenarios returns a set of scenarios to test plaid integration
func getPlaidInvesmentScenarios() ([]plaidInvestmentScenarios, *PlaidWrapper, error) {
	accessToken, err := plaidTestClient.getAccessTokenForSandboxAcct()
	if err != nil {
		return nil, nil, err
	}

	return []plaidInvestmentScenarios{
		{
			// success condition: valid access token
			scenarioName:     "[success condition]: get an account deposits with valid access token",
			shouldErrorOccur: false,
			accessToken:      &accessToken,
		},
		{
			// failure condition: access token is invalid
			scenarioName:     "[failure condition]: get an account deposits with invalid access token",
			shouldErrorOccur: true,
			expectedError:    errors.New("invalid input argument. access token cannot be empty"),
			accessToken:      nil,
		},
	}, plaidTestClient, nil
}

func TestGetPlaidInvestmentsAccountsOperation(t *testing.T) {
	scenarios, p, err := getPlaidInvesmentScenarios()
	assert.Nil(t, err)

	for _, scenario := range scenarios {
		investments, err := p.GetPlaidInvestmentHoldings(context.Background(), scenario.accessToken, nil)
		if err != nil {
			if scenario.shouldErrorOccur {
				assert.Equal(t, err, scenario.expectedError)
			} else {
				t.Errorf("obtained error but not expected - %s", err.Error())
			}
		}

		if scenario.shouldErrorOccur && err == nil {
			t.Errorf("expected error to occur but none did")
		}

		if !scenario.shouldErrorOccur {
			assertAccountsOfInvestmentType(t, investments)
		}
	}
}

func assertAccountsOfInvestmentType(t *testing.T, i *Investments) {
	assert.NotNil(t, i)
	assert.NotNil(t, i.Accounts)
	for _, acct := range i.Accounts {
		assert.NotNil(t, acct)
		assert.Equal(t, plaid.AccountType(INVESTMENTS), acct.Type)
	}
	assert.NotNil(t, i.Securities)
	assert.NotNil(t, i.Holdings)
}

func TestPlaidWrapper_GetPlaidInvestmentHoldings(t *testing.T) {
	type args struct {
		ctx         context.Context
		accessToken *string
		accountIds  []string
	}
	tests := []struct {
		name    string
		p       *PlaidWrapper
		args    args
		want    *Investments
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.GetPlaidInvestmentHoldings(tt.args.ctx, tt.args.accessToken, tt.args.accountIds)
			if (err != nil) != tt.wantErr {
				t.Errorf("PlaidWrapper.GetPlaidInvestmentHoldings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlaidWrapper.GetPlaidInvestmentHoldings() = %v, want %v", got, tt.want)
			}
		})
	}
}
