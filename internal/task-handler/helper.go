package taskhandler

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/sashabaranov/go-openai"
	"go.uber.org/zap"
)

func (th *TaskHandler) callCompletionEndpoint(content *string) (*openai.ChatCompletionResponse, error) {
	var resp openai.ChatCompletionResponse
	var err error
	const maxRetries = 3

	if content == nil {
		return nil, fmt.Errorf("no content")
	}

	// call the openai moderation endpoint to obtain the toxicity score
	req := openai.ChatCompletionRequest{
		Model: openai.GPT4,
		Messages: []openai.ChatCompletionMessage{{
			Role:    openai.ChatMessageRoleUser,
			Content: *content,
		}},
		MaxTokens:        2000,
		Temperature:      0.7,
		TopP:             1.0,
		N:                1,
		Stream:           false,
		Stop:             []string{},
		PresencePenalty:  1,
		FrequencyPenalty: 0.5,
		LogitBias:        map[string]int{},
		User:             "",
		Functions:        []openai.FunctionDefinition{},
		FunctionCall:     nil,
	}

	for i := 0; i < maxRetries; i++ {
		th.logger.Info(fmt.Sprintf("calling completion endpoint. attempt: %d", i+1))
		resp, err = th.openaiClient.CreateChatCompletion(context.Background(), req)
		if err == nil {
			break
		}

		th.logger.Error(fmt.Sprintf("error calling completion endpoint. attempt: %d", i+1), zap.Error(err))

		// Check if it's a rate limit error
		// This could be different based on the API you're using
		// The below check is just an example
		e := &openai.APIError{}
		if errors.As(err, &e) {
			switch e.HTTPStatusCode {
			case 429:
				// rate limiting or engine overload (wait and retry)
				time.Sleep(time.Second * time.Duration(i+1)) // exponential back-off
				continue
			default:
				break
			}
		}
	}

	if err != nil {
		return nil, err
	}

	th.logger.Info("successfully called moderation endpoint")

	return &resp, nil
}
