package handler

import (
	"back_end/model"
	"fmt"
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
		"res": result,
	})

}

// @Summary 测试mysql
// @version 1.0
// @Accept application/x-json-stream
// @Success 200 object model.ResultModel 成功后返回值
// @Router /api/test/insert [POST]

func Insert(ctx *gin.Context) {
	test := model.TestModel{}
	var id = -1
	var message = "数据添加失败"

	if e := ctx.ShouldBind(&test); e == nil {
		id=test.Insert()
		message = "数据添加成功"
	}else {
		fmt.Println(e)
	}
	result := model.ResultModel{
		Code:    http.StatusOK,
		Message: message,
		Data: gin.H{
			"id": id,
		},
	}
	ctx.JSON(http.StatusOK, result)
}

// @Summary 测试mysql
// @version 1.0
// @Accept application/x-json-stream
// @Success 200 object model.ResultModel 成功后返回值
// @Router /api/test/delete [post]