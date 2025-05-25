package esearch

import (
	"github.com/gin-gonic/gin"
)

func NewWebServer(gin *gin.Engine, addr string) {
	err := gin.Run(addr)
	if err != nil {
		panic(err)
	}
}

func NewGrpcWorker() {

}

// NewWebServerTest1 单机模式部署
func NewWebServerTest1(gin *gin.Engine, addr string) {
	err := gin.Run(addr)
	if err != nil {
		panic(err)
	}
}
