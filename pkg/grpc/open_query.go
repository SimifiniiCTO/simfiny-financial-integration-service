package grpc

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

// AskOpenAi is a function that asks OpenAI for information about a particular financial topic
func (s *Server) AskOpenAi(ctx context.Context, prompt string) (*string, error) {
	req := openai.CompletionRequest{
		Model:            openai.GPT4,
		Prompt:           prompt,
		Suffix:           "",
		MaxTokens:        int(s.config.OpenAiMaxTokens),
		Temperature:      s.config.OpenAiTemperature,
		TopP:             s.config.OpenAiTopP,
		N:                0,
		Stream:           false,
		LogProbs:         0,
		Echo:             false,
		Stop:             []string{},
		PresencePenalty:  float32(s.config.OpenAiPresencePenalty),
		FrequencyPenalty: s.config.OpenAiFrequencyPenalty,
		BestOf:           0,
		LogitBias:        map[string]int{},
		User:             "",
	}
	resp, err := s.OpenAiClient.CreateCompletion(ctx, req)
	if err != nil {
		return nil, err
	}

	choices := resp.Choices
	if len(choices) == 0 {
		return nil, fmt.Errorf("invalid response from openai. no choices provided")
	}

	return &choices[0].Text, nil
}
