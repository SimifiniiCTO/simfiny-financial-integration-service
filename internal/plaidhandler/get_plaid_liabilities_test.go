package plaidhandler

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/plaid/plaid-go/v12/plaid"
	"github.com/stretchr/testify/assert"
)

type plaidLiabilityScenarios struct {
	scenarioName     string
	shouldErrorOccur bool
	accessToken      *string
	expectedError    error
}

func getplaidLiabilityScenarios() ([]plaidLiabilityScenarios, *PlaidWrapper, error) {
	accessToken, err := plaidTestClient.getAccessTokenForSandboxAcct()
	if err != nil {
		return nil, nil, err
	}

	return []plaidLiabilityScenarios{
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

func TestGetPlaidAccountLiabilityOperation(t *testing.T) {
	scenarios, p, err := getplaidLiabilityScenarios()
	assert.Nil(t, err)

	for _, scenario := range scenarios {
		liabilities, err := p.getPlaidLiabilities(context.Background(), scenario.accessToken)
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
			assertAccountsOfLiabilityType(t, liabilities)
		}
	}
}

func assertAccountsOfLiabilityType(t *testing.T, d *Liabilities) {
	assert.NotNil(t, d)
	assert.NotNil(t, d.Accounts)
	for _, acct := range d.Accounts {
		assert.NotNil(t, acct)
		if acct.Type != plaid.AccountType(CREDIT) && acct.Type != plaid.AccountType(LOAN) {
			t.Errorf("account can only be of credit or loan type")
		}
	}
}

func TestPlaidWrapper_getPlaidLiabilities(t *testing.T) {
	type args struct {
		ctx         context.Context
		accessToken *string
	}
	tests := []struct {
		name    string
		p       *PlaidWrapper
		args    args
		want    *Liabilities
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.getPlaidLiabilities(tt.args.ctx, tt.args.accessToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("PlaidWrapper.getPlaidLiabilities() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlaidWrapper.getPlaidLiabilities() = %v, want %v", got, tt.want)
			}
		})
	}
}
