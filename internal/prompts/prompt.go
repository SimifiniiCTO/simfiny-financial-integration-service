package prompts

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/pointer"
	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
)

type PromptType string

var DetailedPromptType PromptType = "detailed"
var SummaryPromptType PromptType = "summary"

func trimSpace(v interface{}) (string, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return "", err
	}

	jsonStr := string(bytes)
	noSpaceStr := strings.ReplaceAll(jsonStr, " ", "")

	return noSpaceStr, nil
}

func GenerateInsightsPrompt(financialCtx *schema.MelodyFinancialContext, promptType PromptType) (*string, error) {
	var (
		prompt string
	)

	if financialCtx == nil {
		return nil, errors.New("financial context is nil")
	}

	// promptCtx, err := trimSpace(financialCtx)
	// if err != nil {
	// 	return nil, err
	// }

	if promptType == DetailedPromptType {
		prompt = fmt.Sprintf(
			`act as the best financial copilot available. given %v, 
		 source 10 detailed actionable insights this user can take to 
		 become more financially healthy`, financialCtx)
	} else if promptType == SummaryPromptType {
		prompt = fmt.Sprintf(
			`act as the best financial copilot available. given %v, 
		 source 10 summarized actionable insights to take to  
		 become more financially healthy`, financialCtx)
	}

	return pointer.StringP(prompt), nil
}
