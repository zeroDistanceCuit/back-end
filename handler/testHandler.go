package handler

import (
	"back_end/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary 测试
// @version 1.0
// @Accept application/x-json-stream
// @Success 200 object model.ResultModel 成功后返回值
// @Router /api/test [get]

func Test(ctx *gin.Context) {
	result := model.ResultModel{
		Code:    http.StatusOK,
		Message: "查询成功",
		Data:    "_",
	}
	ctx.JSON(http.StatusOK, gin.H{
		"res":result,
	})

}
