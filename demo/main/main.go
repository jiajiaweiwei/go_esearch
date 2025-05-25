package main

import (
	"demo/handler"
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
	// 启动服务器，监听指定端口
	err := engine.Run("127.0.0.1:" + "23000")
	if err != nil {
		panic(err)
	}
}
