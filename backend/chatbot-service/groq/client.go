package groq

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type chatRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float32   `json:"temperature"`
	MaxTokens   int       `json:"max_tokens"`
}

type chatResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
	Error *struct {
		Message string `json:"message"`
	} `json:"error"`
}

func Ask(systemPrompt, userMessage, model string, temperature float32, maxTokens int) (string, error) {
	apiKey := os.Getenv("GROQ_API_KEY")
	apiURL := os.Getenv("GROQ_API_URL")

	reqBody := chatRequest{
		Model: model,
		Messages: []Message{
			{Role: "system", Content: systemPrompt},
			{Role: "user", Content: userMessage},
		},
		Temperature: temperature,
		MaxTokens:   maxTokens,
	}

	payload, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result chatResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	if result.Error != nil {
		return "", errors.New(result.Error.Message)
	}

	if len(result.Choices) == 0 {
		return "", errors.New("no response from Groq")
	}

	return result.Choices[0].Message.Content, nil
}
