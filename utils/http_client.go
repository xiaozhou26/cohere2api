package utils

import (
	"bytes"
	"cohere/model"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv" // 导入godotenv包
)

func init() {
	// 从当前目录的.env文件加载环境变量
	if err := godotenv.Load(); err != nil {
		fmt.Printf("No .env file found")
	}
}

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
	} else {
		// 从环境变量获取Bearer Token
		tokenEnvVar := os.Getenv("COHERE_BEARER_TOKENS")
		if tokenEnvVar == "" {
			return nil, fmt.Errorf("environment does not contain any authorization tokens")
		}
		tokens := strings.Split(tokenEnvVar, ",")
		if len(tokens) == 0 {
			return nil, fmt.Errorf("no valid tokens found in environment variable")
		}

		selectedToken := tokens[rand.Intn(len(tokens))]
		req.Header.Set("Authorization", selectedToken)
	}
	client := &http.Client{}
	return client.Do(req)
}
