package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"cohere/model"
	"cohere/utils"

	"github.com/gin-gonic/gin"
)

func ChatCompletions(c *gin.Context) {
	var body model.ChatRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	searchQuery := c.Request.URL.Query().Get("q")
	if searchQuery == "" && len(body.Messages) == 0 {
		searchQuery = "hello"
		body.Messages = append(body.Messages, model.Message{Role: "user", Content: searchQuery})
	}

	data := model.ChatData{}
	for _, msg := range body.Messages {
		role := strings.ToUpper(msg.Role)
		if role == "ASSISTANT" {
			role = "CHATBOT"
		}
		data.ChatHistory = append(data.ChatHistory, model.ChatHistory{Role: role, Message: msg.Content})
	}
	data.Message = body.Messages[len(body.Messages)-1].Content
	data.Stream = body.Stream
	data.Model = body.Model

	if strings.HasPrefix(body.Model, "net-") {
		data.Connectors = append(data.Connectors, model.Connector{ID: "web-search"})
	}

	resp, err := utils.FetchChatResponse(data, c.Request.Header.Get("Authorization"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if body.Stream {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Content-Type", "text/event-stream; charset=UTF-8")
		utils.HandleStreamResponse(resp.Body, c.Writer, data.Model)
	} else {
		var result map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, result)
	}
}
