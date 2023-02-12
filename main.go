package main

import (
	"project/bootstrap"
)

func main() {

	// 加载配置文件
	bootstrap.InitializeConfig()
	// 初始化日志
	bootstrap.InitializeLog()
	// 初始化DB
	bootstrap.InitializeDB()
	// 启动服务
	bootstrap.StartServer()
}
