package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func HandleStreamResponse(body io.Reader, w http.ResponseWriter, model string) {
	created := time.Now().Unix()
	scanner := bufio.NewScanner(body)
	for scanner.Scan() {
		var msg map[string]interface{}
		if err := json.Unmarshal(scanner.Bytes(), &msg); err != nil {
			continue
		}

		if text, ok := msg["text"].(string); ok {
			dataMsg := CreateStreamResponse(map[string]interface{}{"text": text}, created, model)
			w.Write([]byte("data: "))
			json.NewEncoder(w).Encode(dataMsg)
			w.Write([]byte("\n\n"))
		}

		if finished, ok := msg["is_finished"].(bool); ok && finished {
			w.Write([]byte("data: [DONE]\n\n"))
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error reading stream: %v", err)
	}
}

func CreateStreamResponse(data map[string]interface{}, created int64, model string) gin.H {
	return gin.H{
		"id":      "chatcmpl-test",
		"object":  "chat.completion.chunk",
		"created": created,
		"model":   model,
		"choices": []gin.H{
			{"index": 0, "delta": gin.H{"role": "assistant", "content": data["text"]}, "finish_reason": nil},
		},
	}
}
