package main

import (
	esearch "framework"
	"framework/demo/handler"

	"github.com/gin-gonic/gin"
)

// 单节点
func main() {
	// 创建默认的 Gin 引擎
	engine := gin.Default()
	// 设置 Gin 运行模式为 Release 模式
	gin.SetMode(gin.ReleaseMode)
	// 设置 POST 请求路由
	engine.POST("/search", handler.SearchTest1)
	// 使用搜索引擎框架
	esearch.NewWebServerTest1(engine, "127.0.0.1:23000")
}
