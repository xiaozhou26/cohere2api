package utils

import (
	"bytes"
	"cohere/model"
	"encoding/json"
	"net/http"
)

func FetchChatResponse(data model.ChatData, authHeader string) (*http.Response, error) {
	url := "https://api.cohere.ai/v1/chat"
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if authHeader != "" {
		req.Header.Set("Authorization", authHeader)
	}
	client := &http.Client{}
	return client.Do(req)
}
