package main

import (
	"cohere/initialize"
	"fmt"
)

func main() {
	// 初始化路由
	router := initialize.InitRouter()

	// 启动服务器
	port := 8080
	router.Run(fmt.Sprintf(":%d", port))
}
