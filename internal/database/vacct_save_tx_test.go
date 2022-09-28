package database

import (
	"context"
	"errors"
	"testing"

	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/helper"
	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/proto"
	"github.com/stretchr/testify/assert"
)

type saveAccountScenario struct {
	scenarioName        string
	shouldErrorOccur    bool
	account             *schema.VirtualAccount
	expectedError       error
	shouldCreateAccount bool
	context             context.Context
}

// saveAccountTestScenarios returns a set of scenarios to test the save account operation
func saveAccountTestScenarios() []saveAccountScenario {
	invalidAccMissingUserID := generateRandomizedAccount()
	invalidAccMissingUserID.UserID = 0

	invalidAccMissingAccessToken := generateRandomizedAccount()
	invalidAccMissingAccessToken.AccessToken = ""
	return []saveAccountScenario{
		{
			// success condition - save a new  account
			scenarioName:        "save new  account",
			shouldErrorOccur:    false,
			account:             generateRandomizedAccount(),
			expectedError:       nil,
			shouldCreateAccount: true,
			context:             context.Background(),
		},
		{
			// failure condition - create a an invalid virtual account (missing necessary fields)
			scenarioName:        "attempt to save an invalid account (account with missing fields) missing userID field",
			shouldErrorOccur:    true,
			account:             invalidAccMissingUserID,
			expectedError:       errors.New("userID cannot be 0"),
			shouldCreateAccount: false,
			context:             context.Background(),
		},
		{
			// failure condition - create a an invalid virtual account (missing necessary fields)
			scenarioName:        "attempt to save an invalid account (account with missing fields) missing accestoken field",
			shouldErrorOccur:    true,
			account:             invalidAccMissingAccessToken,
			expectedError:       errors.New("accesstoken cannot be empty"),
			shouldCreateAccount: false,
			context:             context.Background(),
		},
		{
			// failure condition - save an invalid account
			scenarioName:        "save an invalid account",
			shouldErrorOccur:    true,
			account:             nil,
			expectedError:       errors.New("invalid virtual account parameter. param is nil"),
			shouldCreateAccount: false,
			context:             context.Background(),
		},
	}

}

func TestSaveAccountOperation(t *testing.T) {
	scenarios := saveAccountTestScenarios()
	reset()

	for _, scenario := range scenarios {
		var createdAcct *schema.VirtualAccount = scenario.account
		var err error
		if scenario.shouldCreateAccount {
			accesstoken := helper.GenerateRandomString(10)
			createdAcct, err = Conn.CreateVirtualAccount(scenario.context, scenario.account, accesstoken)
			if err != nil {
				t.Errorf("failed to create test account as precondition - %s", err.Error())
			}

			assert.NotNil(t, createdAcct)

			{
				// update the access token and set acct as inactive
				createdAcct.AccessToken = helper.GenerateRandomString(10)
			}
		}

		if err := Conn.SaveVirtualAccountRecord(scenario.context, createdAcct); err != nil {
			if scenario.shouldErrorOccur {
				assert.Equal(t, scenario.expectedError, err)
			} else {
				t.Errorf("obtained error but not expected - %s", err.Error())
			}
		} else {
			if scenario.shouldErrorOccur {
				t.Errorf("expected error %s but none occured", scenario.expectedError.Error())
			}
		}

		if !scenario.shouldErrorOccur {
			// attempt to get the account from the database
			if createdAcct != nil {
				acct, err := Conn.FindVirtualAcctID(scenario.context, createdAcct.Id)
				if err != nil {
					t.Errorf("obtained error but not expected - %s", err.Error())
				}

				assert.NotNil(t, acct)
			}
		}
	}
}
