package initialize

import (
	"cohere/handler"

	"github.com/gin-gonic/gin"
)

// InitRouter initializes the Gin router with all necessary routes and middleware
func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(handler.CORSMiddleware())

	r.POST("/v1/chat/completions", handler.ChatCompletions)
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})
	return r
}
