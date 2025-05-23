package main

import (
	"context"
	"fmt"
	"log"

	openai "github.com/sashabaranov/go-openai"
)

type DeepSeekMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type DeepSeekRequest struct {
	Model       string            `json:"model"`
	Messages    []DeepSeekMessage `json:"messages"`
	Temperature float32           `json:"temperature,omitempty"`
	Stream      bool              `json:"stream,omitempty"`
}

type DeepSeekResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

func generateWithDeepSeek(diff string, config *Config) (string, error) {
	apiKey, model, temperature, baseURL := config.GetActiveConfig()

	log.Printf("Using DeepSeek config - Model: %s, BaseURL: %s", model, baseURL)

	if baseURL == "" {
		baseURL = "https://api.deepseek.com"
	}

	clientConfig := openai.DefaultConfig(apiKey)
	clientConfig.BaseURL = baseURL

	client := openai.NewClientWithConfig(clientConfig)

	prompt := buildPrompt(diff)
	log.Printf("Sending prompt to DeepSeek API, length: %d", len(prompt))

	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: prompt,
		},
	}

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:       model,
			Messages:    messages,
			Temperature: temperature,
		},
	)

	if err != nil {
		log.Printf("DeepSeek API error: %v", err)
		return "", fmt.Errorf("failed to generate commit message with DeepSeek: %w", err)
	}

	if len(resp.Choices) == 0 {
		log.Println("No choices in DeepSeek API response")
		return "", fmt.Errorf("no response from DeepSeek API")
	}

	message := resp.Choices[0].Message.Content
	log.Printf("Successfully generated message, length: %d", len(message))
	return processMessage(message)
}
