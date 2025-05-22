package main

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

func generateWithOpenAI(diff string, config *Config) (string, error) {
	apiKey, model, temperature, _ := config.GetActiveConfig()

	client := openai.NewClient(apiKey)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: buildPrompt(diff),
				},
			},
			Temperature: temperature,
		},
	)

	if err != nil {
		return "", fmt.Errorf("failed to generate commit message: %w", err)
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no response from OpenAI")
	}

	message := resp.Choices[0].Message.Content
	return processMessage(message)
}
