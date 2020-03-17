package handler

import (
	"back_end/config/authConfig"
	"back_end/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

func BussinessRegister(ctx *gin.Context)  {
	user:=model.BussinessModel{}
	var id=-1
	var message="用户注册失败,程序猿会马上修补bug"
	var code=http.StatusOK

	if e := ctx.BindJSON(&user); e != nil {
		message = "数据绑定失败"
		code = http.StatusUnauthorized
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"result": "error",
		})
	}

	u:= user.QueryByUsername();

	if u.Name != ""{
		message="该账号已存在，请另起名字"
	}else{
		id=user.Insert()
		message="用户"+user.Name+",注册成功"
	}

	result:=model.ResultModel{
		Code:    code,
		Message: message,
		Data:    gin.H{
			"id":id,
		},
	}

	ctx.JSON(http.StatusOK,result)
}

func BussinessCreateJwt(ctx *gin.Context)  {
	// 获取用户
	user := &model.BussinessModel{}
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
	if u.Name!="" && u.Password == user.Password {
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
				"userId":u.Id,
			})
		} else {
			result.Message = "登录失败"
			result.Code = http.StatusBadRequest
			ctx.JSON(result.Code, gin.H{
				"result": result,
			})
		}
	} else {
		result.Message = "登录失败"
		result.Code = http.StatusBadRequest
		ctx.JSON(result.Code, gin.H{
			"result": result,
		})
	}
}

// 修改密码
func BussinessUpdatePassW(ctx *gin.Context)  {
	result := &model.ResultModel{
		Code:    200,
		Message: "修改成功",
		Data:    nil,
	}
	id := ctx.Param("id")

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
	user := model.BussinessModel{
		Id: i,
	}
	if e := ctx.BindJSON(&user); e != nil {
		result.Message = "数据绑定失败"
		result.Code = http.StatusUnauthorized
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"result": result,
		})
	}

	flag:=user.UpdatePassW()

	if !flag {
		result.Message="修改失败"
	}
	ctx.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

//查询用户信息
func GetOneBussiness(ctx *gin.Context)  {
	//重构检测id字符型
	id := ctx.Param("id")
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

	user:=model.BussinessModel{
		Id:       i,
	}

	userInfo:=user.GetOneBussinessInfo()

	userInfo.Password="****"

	ctx.JSON(http.StatusOK, gin.H{
		"result": model.ResultModel{
			Code:    http.StatusOK,
			Message: "查询成功",
			Data:    userInfo,
		},
	})
}