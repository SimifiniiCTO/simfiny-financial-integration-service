package grpc

import (
	"context"
	"errors"
	"testing"

	"github.com/SimifiniiCTO/simfinii/src/backend/services/financial-integration-service/internal/helper"
	proto "github.com/SimifiniiCTO/simfinii/src/backend/services/financial-integration-service/proto"
	"github.com/stretchr/testify/assert"
)

type GetLinkTokenRpcScenario struct {
	scenarioName     string
	shouldErrorOccur bool
	expectedError    error
	userID           uint64
}

func GetLinkTokenRpcScenarios() ([]GetLinkTokenRpcScenario, error) {
	return []GetLinkTokenRpcScenario{
		{
			scenarioName:     "[success scenario] - get an account that already exists",
			shouldErrorOccur: false,
			expectedError:    nil,
			userID:           uint64(helper.GenerateRandomId(10000, 20000)),
		},
		{
			scenarioName:     "[failure scenario] - get an account that doesn't already exist",
			shouldErrorOccur: true,
			expectedError:    errors.New("invalid request object. user id cannot be 0"),
			userID:           0,
		},
	}, nil
}

func TestGetLinkTokenRpcOperation(t *testing.T) {
	ctx := context.Background()
	conn := MockGRPCService(ctx)
	defer conn.Close()

	scenarios, err := GetLinkTokenRpcScenarios()
	assert.Nil(t, err)

	for _, scenario := range scenarios {
		c := proto.NewFinancialIntegrationServiceClient(conn)

		req := &proto.InitiateAccountSetupRequest{
			UserID: scenario.userID,
		}

		linkToken, err := c.InitiateAccountSetupAndGetLinkToken(ctx, req)
		if err != nil {
			if scenario.shouldErrorOccur {
				assert.Contains(t, err.Error(), scenario.expectedError.Error())
			} else {
				t.Errorf("obtained error but not expected - %s", err.Error())
			}
		} else {
			if scenario.shouldErrorOccur {
				t.Errorf("expected error %s but none occured", scenario.expectedError.Error())
			}
		}

		if !scenario.shouldErrorOccur {
			assert.NotNil(t, linkToken)
		}
	}
}
