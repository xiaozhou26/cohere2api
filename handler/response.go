package handler

import (
	"github.com/gin-gonic/gin"
)

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
