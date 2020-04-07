package handler

import (
	"back_end/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetUserFleshModel(ctx *gin.Context){
	id := ctx.Query("id")
	i, e := strconv.Atoi(id)
	if e != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"result": model.ResultModel{
				Code:    http.StatusBadRequest,
				Message: "id 不是 int 类型, id 转换失败",
				Data:    e.Error(),
			},
		})
		log.Panicln("id 不是 int 类型, id 转换失败", e.Error())
	}

	user := model.UserFleshModel{
		Id: i,
	}

	userFlesh:=user.Find()

	ctx.JSON(http.StatusOK, gin.H{
		"result": model.ResultModel{
			Code:    http.StatusOK,
			Message: "查询成功",
			Data:    userFlesh,
		},
	})
}
