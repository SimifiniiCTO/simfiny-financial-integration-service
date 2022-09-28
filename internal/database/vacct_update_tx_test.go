package database

import (
	"context"
	"errors"
	"testing"

	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/helper"
	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/proto"
	"github.com/stretchr/testify/assert"
)

type updateAccountScenario struct {
	scenarioName        string
	account             *schema.VirtualAccount
	shouldCreateAccount bool
	shouldErrorOccur    bool
	expectedError       error
	context             context.Context
}

func updateAccountScenarios() []updateAccountScenario {
	nonExistingVAcct := generateRandomizedAccount()
	nonExistingVAcct.Id = uint64(helper.GenerateRandomId(3000, 400000))
	return []updateAccountScenario{
		{
			// success condition - update existing merchant account
			scenarioName:        "update existing account",
			account:             generateRandomizedAccount(),
			shouldCreateAccount: true,
			shouldErrorOccur:    false,
			expectedError:       nil,
			context:             context.Background(),
		},
		{
			// failure condition - update a non-existing merchant account
			scenarioName:        "update non-existing account",
			account:             &schema.VirtualAccount{},
			shouldCreateAccount: false,
			shouldErrorOccur:    true,
			expectedError:       errors.New("invalid input argument. virtual account id cannot be 0"),
			context:             context.Background(),
		},
		{
			// failure condition - update a non-existing merchant account
			scenarioName:        "update non-existing account account not nil",
			account:             nonExistingVAcct,
			shouldCreateAccount: false,
			shouldErrorOccur:    true,
			expectedError:       errors.New("record not found"),
			context:             context.Background(),
		},
	}
}

func TestUpdateAccountOperation(t *testing.T) {
	scenarios := updateAccountScenarios()
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
		}

		vAcct.UserID = uint64(helper.GenerateRandomId(10000, 3000000))
		if err = Conn.UpdateVirtualAccount(scenario.context, vAcct, vAcct.Id); err != nil {
			if scenario.shouldErrorOccur {
				assert.Equal(t, err, scenario.expectedError)
			} else {
				t.Errorf("obtained error but not expected - %s", err.Error())
			}
		}

		if scenario.shouldErrorOccur && err == nil {
			t.Errorf("expected error to occur but none did")
		}
	}
}
