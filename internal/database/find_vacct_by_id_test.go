package database

import (
	"context"
	"errors"
	"testing"

	"github.com/SimifiniiCTO/simfinii/src/backend/services/financial-integration-service/internal/helper"
	schema "github.com/SimifiniiCTO/simfinii/src/backend/services/financial-integration-service/proto"
	"github.com/stretchr/testify/assert"
)

type findAccountByIdScenario struct {
	scenarioName        string
	shouldErrorOccur    bool
	account             *schema.VirtualAccount
	shouldCreateAccount bool
	expectedError       error
	deactivateAccount   bool
	context             context.Context
}

// findAccountByIdScenarios returns a set of scenarios to test the account's existence based on provided email
func findAccountByIdScenarios() []findAccountByIdScenario {
	acctWithNonExistendVacct := generateRandomizedAccount()
	acctWithNonExistendVacct.Id = uint64(helper.GenerateRandomId(10000, 200000))
	return []findAccountByIdScenario{
		{
			// success condition: account exists
			scenarioName:        "find an account by virtual account id that already exists",
			shouldErrorOccur:    false,
			account:             generateRandomizedAccount(),
			shouldCreateAccount: true,
			expectedError:       nil,
			context:             context.Background(),
		},
		{
			// failure condition: account does not exist and vAcct id is 0
			scenarioName:        "[failure scenario]: attempt to find an account by invalid user acct id",
			shouldErrorOccur:    true,
			account:             generateRandomizedAccount(),
			shouldCreateAccount: false,
			expectedError:       errors.New("invalid input argument. virtual account cannot be 0"),
			context:             context.Background(),
		},
		{
			// failure condition: account does not exist
			scenarioName:        "[failure scenario]: attempt to find an account by a virtual acct id that does not exist",
			shouldErrorOccur:    true,
			account:             acctWithNonExistendVacct,
			shouldCreateAccount: false,
			expectedError:       errors.New("record not found"),
			context:             context.Background(),
		},
		{
			// failure condition: account exists but was soft deleted
			scenarioName:        "[failure scenario]: attempt to find an account that was soft deleted by user acct id",
			shouldErrorOccur:    true,
			account:             generateRandomizedAccount(),
			shouldCreateAccount: true,
			expectedError:       errors.New("virtual account does not exist. account was soft deleted"),
			deactivateAccount:   true,
			context:             context.Background(),
		},
	}
}

func TestFindAccountByIdOperation(t *testing.T) {
	scenarios := findAccountByIdScenarios()
	reset()

	for _, scenario := range scenarios {
		var vAcct = scenario.account
		var err error

		if scenario.shouldCreateAccount {
			accessToken := helper.GenerateRandomString(10)
			vAcct, err = Conn.CreateVirtualAccount(scenario.context, scenario.account, accessToken)
			if err != nil {
				t.Errorf("obtained error but not expected - %s", err.Error())
			}

			if scenario.deactivateAccount {
				if err := Conn.DeactivateVirtualAccount(scenario.context, vAcct.Id); err != nil {
					t.Errorf("obtained error but not expected - %s", err.Error())
				}
			}
		}

		acct, err := Conn.FindVirtualAcctID(scenario.context, vAcct.Id)
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
			assert.NotNil(t, acct)
		}
	}
}
