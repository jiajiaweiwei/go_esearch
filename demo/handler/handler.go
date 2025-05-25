package handler

import (
	"demo/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SearchTest1(ctx *gin.Context) {
	var helloReq types.HelloSearchReq

	err := ctx.ShouldBindJSON(&helloReq)
	if err != nil {
		result := make(map[string]string)
		result["result"] = "Fail."
		ctx.JSON(http.StatusInternalServerError, result)
		return
	}
	ctx.JSON(http.StatusOK, helloReq)
	return
}
