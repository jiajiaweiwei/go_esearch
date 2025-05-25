package handler

import (
	"framework/demo/searcher"
	"framework/demo/types"
	index_service "framework/frame/server/worker/index"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 实例化搜索引擎
var Indexer index_service.Indexer

func init() {
	Indexer = index_service.NewSentinel([]string{"127.0.0.1:2379"})
}

func SearchTest1(ctx *gin.Context) {
	// 1.解析请求数据
	var helloReq types.HelloSearchReq
	err := ctx.ShouldBindJSON(&helloReq)
	if err != nil {
		result := make(map[string]string)
		result["result"] = "Fail."
		ctx.JSON(http.StatusInternalServerError, result)
		return
	}

	// 2.根据需求获取搜索器
	searcherByMod1 := searcher.NewStringSearcherByMod1()
	responseString := searcherByMod1.Search(&helloReq, Indexer)

	// 3.执行搜索，返回结果
	ctx.JSON(http.StatusOK, responseString)
	return
}
