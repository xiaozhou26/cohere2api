package main

import (
	"cohere/initialize"
	"fmt"
	"os"

	"github.com/joho/godotenv" // 导入godotenv包
)

var (
	Port string
	Bind string
)

func main() {
	// 从当前目录的.env文件加载环境变量
	godotenv.Load()
	// 初始化路由
	router := initialize.InitRouter()
	// 启动服务器
	Port = os.Getenv("PORT")
	if Port == "" {
		Port = "8080"
	}
	if Bind == "" {
		Bind = "0.0.0.0"
	}
	router.Run(fmt.Sprint(Bind, ":", Port))
}
