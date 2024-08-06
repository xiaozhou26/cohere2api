package handler

import (
	"cohere/utils"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetModelsHandler(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")

	resp, err := utils.FetchModelInfo(authHeader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	c.Header("Content-Type", resp.Header.Get("Content-Type"))
	c.Status(resp.StatusCode)

	_, err = io.Copy(c.Writer, resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to stream response"})
		return
	}
}
