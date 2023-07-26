package taskhandler

import (
	"context"
	"encoding/json"

	"github.com/SimifiniiCTO/asynq"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/pointer"
	"go.uber.org/zap"
)

type GenerateActionableInsightPayload struct{}

func (a GenerateActionableInsightPayload) String() *string {
	return pointer.StringP("GenerateActionableInsightPayload")
}

// NewGenerateActionableInsights initialize an actionable insights task
func NewGenerateActionableInsights() (*asynq.Task, error) {
	payload, err := json.Marshal(&GenerateActionableInsightPayload{})
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TaskGenerateActionableInsights.String(), payload), nil
}

func (th *TaskHandler) RunGenerateActionableInsights(ctx context.Context, task *asynq.Task) error {
	var (
		payload GenerateActionableInsightPayload
		db      = th.clickhouseDb
	)

	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return err
	}

	// get all user profiles on the platform
	profiles, err := th.postgresDb.GetAllUserProfiles(ctx)
	if err != nil {
		return err
	}

	// for each profile, query and obtain its financial context
	// only if it has a connected plaid account with us
	for _, profile := range profiles {
		if len(profile.Link) > 0 {
			// retrieve the financial context for the user
			userId := profile.UserId
			// TODO: optimize this to perform this in one single query and obtain all financial contexts on the platform
			financialContext, err := db.GetFinancialContextForCurrentMonth(ctx, &userId, 5)
			if err != nil {
				// log the error
				th.logger.Error("failed to get financial context for user", zap.Any("userId", userId), zap.Error(err))
				continue
			}

			// call open ai with the financial context

		}
	}
}
