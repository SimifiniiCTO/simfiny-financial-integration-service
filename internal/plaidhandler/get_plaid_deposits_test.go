package plaidhandler

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/plaid/plaid-go/v12/plaid"
	"github.com/stretchr/testify/assert"
)

type plaidDepositScenarios struct {
	scenarioName     string
	shouldErrorOccur bool
	accessToken      *string
	expectedError    error
}

// findAccountByIdScenarios returns a set of scenarios to test the account's existence based on provided email
func getPlaidDepositScenarios() ([]plaidDepositScenarios, *PlaidWrapper, error) {
	accessToken, err := plaidTestClient.getAccessTokenForSandboxAcct()
	if err != nil {
		return nil, nil, err
	}

	return []plaidDepositScenarios{
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

func TestGetPlaidDepositoryAccountsOperation(t *testing.T) {
	scenarios, p, err := getPlaidDepositScenarios()
	assert.Nil(t, err)

	for _, scenario := range scenarios {
		deposits, err := p.getPlaidDeposit(context.Background(), scenario.accessToken)
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
			assertAccountsOfDepositoryType(t, deposits)
		}
	}
}

func assertAccountsOfDepositoryType(t *testing.T, d *Deposits) {
	assert.NotNil(t, d)
	for _, acct := range d.Accounts {
		assert.NotNil(t, acct)
		assert.Equal(t, plaid.AccountType(DEPOSITORY), acct.Type)
	}
}

func TestPlaidWrapper_getPlaidDeposit(t *testing.T) {
	type args struct {
		ctx         context.Context
		accessToken *string
	}
	tests := []struct {
		name    string
		p       *PlaidWrapper
		args    args
		want    *Deposits
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.getPlaidDeposit(tt.args.ctx, tt.args.accessToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("PlaidWrapper.getPlaidDeposit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlaidWrapper.getPlaidDeposit() = %v, want %v", got, tt.want)
			}
		})
	}
}
