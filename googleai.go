package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type GoogleAIMessage struct {
	Content string `json:"content"`
}

type GoogleAIContent struct {
	Parts []struct {
		Text string `json:"text"`
	} `json:"parts"`
}

type GoogleAIRequest struct {
	Contents []struct {
		Parts []struct {
			Text string `json:"text"`
		} `json:"parts"`
	} `json:"contents"`
}

type GoogleAIResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}

func generateWithGoogleAI(diff string, config *Config) (string, error) {
	apiKey, model, _, _ := config.GetActiveConfig()

	if apiKey == "" {
		return "", fmt.Errorf("Google AI API key is not set")
	}

	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/%s:generateContent?key=%s", model, apiKey)

	messages := []GoogleAIMessage{
		{
			Content: buildPrompt(diff),
		},
	}

	reqBody := GoogleAIRequest{
		Contents: []struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		}{
			{
				Parts: []struct {
					Text string `json:"text"`
				}{
					{Text: messages[0].Content},
				},
			},
		},
	}

	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("API request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API request failed with status: %s", resp.Status)
	}

	var response GoogleAIResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	if len(response.Candidates) == 0 || len(response.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("no content in response")
	}

	return response.Candidates[0].Content.Parts[0].Text, nil
}
