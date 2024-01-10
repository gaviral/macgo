// File: openaiapi/client.go

package openaiapi

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const openAIImageEndpoint = "https://api.openai.com/v1/images/generations"

type ImageRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	N      int    `json:"n"`
	Size   string `json:"size"`
}

type ImageResponse struct {
	Data []struct {
		URL string `json:"url"`
	} `json:"data"`
}

func authenticateRequest(req *http.Request) {
	const openAIKey = ""
	req.Header.Set("Authorization", "Bearer "+openAIKey)
}

func CreateImage(prompt string, n int, size string) (*ImageResponse, error) {
	payload := ImageRequest{
		Model:  "dall-e-3",
		Prompt: prompt,
		N:      n,
		Size:   size,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", openAIImageEndpoint, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, err
	}

	authenticateRequest(req)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result ImageResponse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
