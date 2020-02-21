package handler

import (
	"back_end/model"
	"back_end/config/authConfig"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"time"
)

func Register(ctx *gin.Context)  {
	user:=model.UserModel{}
	var id=-1
	var message="用户注册失败"

	if e:=ctx.BindJSON(&user);e==nil{
		fmt.Println(user)
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

func CreateJwt(ctx *gin.Context)  {
	// 获取用户
	user := &model.UserModel{}
	result := &model.ResultModel{
		Code:    200,
		Message: "登录成功",
		Data:    nil,
	}


	if e := ctx.BindJSON(&user); e != nil {
		result.Message = "数据绑定失败"
		result.Code = http.StatusUnauthorized
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"result": result,
		})
	}

	u:= user.QueryByUsername()

	if u.Password == user.Password {
		expiresTime := time.Now().Unix() + int64(authConfig.OneDayOfHours)
		claims := jwt.StandardClaims{
			Audience:  user.Name,     // 受众
			ExpiresAt: expiresTime,       // 失效时间
			Id:        string(user.Id),   // 编号
			IssuedAt:  time.Now().Unix(), // 签发时间
			Issuer:    "gin hello",       // 签发人
			NotBefore: time.Now().Unix(), // 生效时间
			Subject:   "login",           // 主题
		}
		var jwtSecret = []byte(authConfig.Secret)
		tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		if token, err := tokenClaims.SignedString(jwtSecret); err == nil {
			result.Message = "登录成功"
			result.Data = "Bearer " + token
			result.Code = http.StatusOK
			ctx.JSON(result.Code, gin.H{
				"result": result,
			})
		} else {
			result.Message = "登录失败"
			result.Code = http.StatusOK
			ctx.JSON(result.Code, gin.H{
				"result": result,
			})
		}
	} else {
		result.Message = "登录失败"
		result.Code = http.StatusOK
		ctx.JSON(result.Code, gin.H{
			"result": result,
		})
	}
}
