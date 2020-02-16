package handler

import (
	"back_end/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// @Summary 测试
// @version 1.0
// @Tags test
// @Accept application/x-json-stream
// @Success 200 object model.ResultModel 成功后返回值
// @Failure 409 object model.Result 添加失败
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

// @Summary 测试mysql-查询所有数据
// @version 1.0
// @Tags test
// @Accept application/x-json-stream
// @Success 200 object model.ResultModel 成功后返回值
// @Router /api/test/findAll [get]
func GetAll(ctx *gin.Context) {
	test := model.TestModel{}
	testRes := test.FindAll()
	result := model.ResultModel{
		Code:    http.StatusOK,
		Message: "查询成功",
		Data:    testRes,
	}
	ctx.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

// @Summary 测试mysql-增加数据
// @version 1.0
// @Id 1
// @Tags test
// @Accept application/x-json-stream
// @Success 200 object model.ResultModel 成功后返回值
// @Failure 409 object model.Result 添加失败
// @Router /api/test/insert [post]

func Insert(ctx *gin.Context) {
	test := model.TestModel{}
	var id = -1
	var message = "数据添加失败"

	if e := ctx.ShouldBind(&test); e == nil {
		id = test.Insert()
		message = "数据添加成功"
	} else {
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

// @Summary 测试mysql-删除
// @version 1.0
// @Tags test
// @Accept application/x-json-stream
// @Success 200 object model.ResultMo
//del 成功后返回值
// @Router /api/test/save [delete]

func Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	i, e := strconv.Atoi(id)
	if e != nil {
		log.Panicln("id 不是 int 类型, id 转换失败", e.Error())
	}
	test := model.TestModel{Id: i}
	test.Delete()
}

// @Summary 测试mysql-查询单个
// @version 1.0
// @Tags test
// @Accept application/x-json-stream
// @Success 200 object model.ResultModel 成功后返回值
// @Router /api/test/getOne [get]

func GetOne(ctx *gin.Context) {
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

	test := model.TestModel{
		Id: i,
	}

	testOne := test.FindById()
	ctx.JSON(http.StatusOK, gin.H{
		"result": model.ResultModel{
			Code:    http.StatusOK,
			Message: "查询成功",
			Data:    testOne,
		},
	})
}

// @Summary 测试mysql-修改
// @version 1.0
// @Tags test
// @Accept application/x-json-stream
// @Success 200 object model.ResultModel 成功后返回值
// @Router /api/test/getOne [post]

func Update(ctx *gin.Context) {
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
	test := model.TestModel{
		Id: i,
	}
	var message = "数据更新失败"

	if e := ctx.ShouldBind(&test); e == nil {
		test.SaveHandler()
		message = "数据更新成功"
	} else {
		fmt.Println(e)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"result": model.ResultModel{
			Code:    http.StatusOK,
			Message: message,
			Data:    test.Id,
		},
	})
}
