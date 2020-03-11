package handler

import (
	"back_end/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

//获取商品信息
func ShopSearch(ctx *gin.Context){
	goods:=&model.GoodsModel{}
	result:=&model.ResultModel{
		Code:    200,
		Message: "查询失败",
		Data:    nil,
	}

	goods.Name=ctx.DefaultQuery("name","goods")
	goods.Type=ctx.Query("type")
	goodsArr:=goods.FindByName()

	// 查询类型还没有加入
	//TODO 加入正则表达式，正则搜索相关内容
	if len(goodsArr)>= 1 {
		result.Message="查询成功"
		result.Data=goodsArr
	}
	ctx.JSON(http.StatusOK,gin.H{
		"result":result,
	})
}
