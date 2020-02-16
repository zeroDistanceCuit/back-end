package handler

import (
	"back_end/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Register(ctx *gin.Context)  {
	user:=model.UserModel{}
	var id=-1
	var message="用户注册失败"

	if e:=ctx.ShouldBind(&user);e==nil{
		id=user.Insert()
		message="用户"+user.Name+",注册成功"
	}else {
		log.Panicln(e)
	}

	result:=model.ResultModel{
		Code:    http.StatusOK,
		Message: message,
		Data:    gin.H{
			"id":id,
		},
	}
	ctx.JSON(http.StatusOK,result)
}
