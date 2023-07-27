package taskhandler

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/SimifiniiCTO/asynq"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/pointer"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/prompts"
	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"github.com/sashabaranov/go-openai"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	financialContexts, err := th.clickhouseDb.GetAllFinancialContextsForCurrentMonth(ctx, 4)
	if err != nil {
		return err
	}

	for userId, financialContext := range financialContexts {
		// call openAI for detailed response
		detailedResponse, err := th.getActionableInsight(financialContext, prompts.DetailedPromptType)
		if err != nil {
			th.logger.Error(err.Error())
			continue
		}

		summarizedResponse, err := th.getActionableInsight(financialContext, prompts.SummaryPromptType)
		if err != nil {
			th.logger.Error(err.Error())
			continue
		}

		insight := &schema.ActionableInsight{
			Id:               userId,
			DetailedAction:   detailedResponse.Content,
			SummarizedAction: summarizedResponse.Content,
			GeneratedTime:    timestamppb.New(time.Now()),
			Tags:             []string{},
		}

		if err := th.postgresDb.AddActionableInsight(ctx, userId, insight); err != nil {
			th.logger.Error(err.Error())
			continue
		}
	}

	return nil
}

func (th *TaskHandler) getActionableInsight(financialContext *schema.MelodyFinancialContext, promptType prompts.PromptType) (*openai.ChatCompletionMessage, error) {
	detailedPromt, err := prompts.GenerateInsightsPrompt(financialContext, promptType)
	if err != nil {
		th.logger.Error("Failed to generate financial context prompt for user", zap.Error(err))
		return nil, err
	}

	response, err := th.callCompletionEndpoint(detailedPromt)
	if err != nil {
		return nil, err
	}

	if len(response.Choices) > 0 {
		return &response.Choices[0].Message, nil
	}

	return nil, fmt.Errorf("Failed to generate a response from open-ai using the financial context")
}
