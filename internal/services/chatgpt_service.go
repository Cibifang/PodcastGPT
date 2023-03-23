package services

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type ChatGPTService struct {
	Endpoint string
}

type ChatGPTResponse struct {
	Response string `json:"response"`
	Err      string `json:"err"`
}

func NewChatGPTService() *ChatGPTService {
	return &ChatGPTService{
		Endpoint: "https://api.openai.com/v1/engines/davinci-codex/completions",
	}
}

func (c *ChatGPTService) GetChatResponse(prompt string) (string, error) {
	type ChatGPTRequest struct {
		Prompt      string   `json:"prompt"`
		MaxTokens   int      `json:"max_tokens"`
		Temperature float32  `json:"temperature"`
		Stop        []string `json:"stop"`
	}

	reqBody, err := json.Marshal(&ChatGPTRequest{
		Prompt:      prompt,
		MaxTokens:   100,
		Temperature: 0.7,
		Stop:        []string{"\n"},
	})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, c.Endpoint, bytes.NewReader(reqBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer YOUR_OPENAI_API_KEY")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var chatResp ChatGPTResponse
	if err := json.Unmarshal(respBody, &chatResp); err != nil {
		return "", err
	}

	if chatResp.Err != "" {
		return "", errors.New(chatResp.Err)
	}

	return strings.TrimSpace(chatResp.Response), nil
}