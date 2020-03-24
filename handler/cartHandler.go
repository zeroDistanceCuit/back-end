package handler

import (
	"back_end/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// 加入购物车
func InsertNewShop(ctx *gin.Context){
	cart := &model.CartModel{}
	result := &model.ResultModel{
		Code:    200,
		Message: "新增成功",
		Data:    nil,
	}

	if e := ctx.BindJSON(&cart); e != nil {
		result.Message = "数据绑定失败"
		result.Code = http.StatusUnauthorized
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"result": result,
		})
	}

	// 将商家库存减少[x],成交时减少，避免返回商
	cart.Insert()

	ctx.JSON(http.StatusOK,gin.H{
		"result":result,
	})
}

// 查询自己的购物车
func GetShopCartList(ctx *gin.Context){
	cart := &model.CartModel{}
	result := &model.ResultModel{
		Code:    200,
		Message: "查询成功",
		Data:    nil,
	}

	userId:=ctx.Query("userId")

	i, e := strconv.Atoi(userId)
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

	cart.UserId=i
	cart.Status=ctx.Query("status")
	cartList:=cart.SearchByUserId()
	result.Data=cartList
	ctx.JSON(http.StatusOK,gin.H{
		"result":result,
	})
}

//付款后修改商品状态
func Payment(ctx *gin.Context)  {
	cart := &model.CartModel{}
	shop:=&model.ShopsModel{}
	result := &model.ResultModel{
		Code:    200,
		Message: "付款成功",
		Data:    nil,
	}

	if e := ctx.BindJSON(&cart); e != nil {
		result.Message = "数据绑定失败"
		result.Code = http.StatusUnauthorized
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"result": result,
		})
	}

	shop.Id=cart.ShopsId
	shopEx:=shop.GetById()
	shopEx.Nums-=cart.Num
	shopEx.Update()

	status:=cart.UpdateByOrder()

	if !status{
		result.Message = "付款失败"
		result.Code = http.StatusUnauthorized
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"result": result,
		})
	}else{
		ctx.JSON(http.StatusOK,gin.H{
			"result":result,
		})
	}
}

//移除购物车
func RemoveCart(ctx *gin.Context)  {
	cart := &model.CartModel{}
	result := &model.ResultModel{
		Code:    200,
		Message: "移除成功",
		Data:    nil,
	}

	if e := ctx.BindJSON(&cart); e != nil {
		result.Message = "数据绑定失败"
		result.Code = http.StatusUnauthorized
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"result": result,
		})
	}

	if flag:=cart.Delete();flag == false{
		result.Message = "移除失败"
		result.Code = http.StatusUnauthorized
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"result": result,
		})
	}else{
		ctx.JSON(http.StatusOK,gin.H{
			"result":result,
		})
	}
}