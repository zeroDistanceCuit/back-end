package handler

import (
	"back_end/config/initDB"
	"back_end/model"
	"github.com/gin-gonic/gin"
	_ "gopkg.in/go-playground/validator.v9"
	"log"
	"net/http"
	"strconv"
)

//获取商品信息
func ShopSearch(ctx *gin.Context) {
	goods := &model.GoodsModel{}
	result := &model.ResultModel{
		Code:    200,
		Message: "查询失败",
		Data:    nil,
	}

	goods.Name = ctx.DefaultQuery("name", "goods")
	goods.Type = ctx.Query("type")
	goodsArr := goods.FindByName()

	// 查询类型还没有加入
	//TODO 加入正则表达式，正则搜索相关内容
	if len(goodsArr) >= 1 {
		result.Message = "查询成功"
		result.Data = goodsArr
	}
	ctx.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

//给店铺增加商品
//先查出来，有则加无则创建
//需要考虑是否含有本条记录
func AddShops(ctx *gin.Context) {
	shops := &model.ShopsModel{}
	result := &model.ResultModel{
		Code:    200,
		Message: "新增失败",
		Data:    nil,
	}

	if e := ctx.BindJSON(&shops); e != nil {
		result.Message = "数据绑定失败"
		result.Code = http.StatusUnauthorized
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"result": result,
		})
	}

	num := shops.Nums

	shopEx := shops.Search()

	shopEx.Nums += num

	flag := shopEx.Update()

	if flag {
		result.Message = "新增成功"
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

//测试
//减少商品
//先查出来，有则加无则创建
//需要考虑是否含有本条记录
func UpdateShops(ctx *gin.Context) {
	shops := &model.ShopsModel{}
	result := &model.ResultModel{
		Code:    200,
		Message: "修改失败",
		Data:    nil,
	}

	if e := ctx.BindJSON(&shops); e != nil {
		result.Message = "数据绑定失败"
		result.Code = http.StatusUnauthorized
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"result": result,
		})
	}
	num := shops.Nums

	shopEx := shops.Search()

	shopEx.Nums = num

	flag := shopEx.Update()

	if flag {
		result.Message = "修改成功"
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

// 关联查询
func SearchAllShops(ctx *gin.Context) {
	shopsInfo := &model.ShopsModel{}
	result := &model.ResultModel{
		Code:    200,
		Message: "查询成功",
		Data:    nil,
	}

	bussinessId := ctx.Query("bussinessId")
	i, e := strconv.Atoi(bussinessId)
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

	shopsInfo.BussinessId = i

	shopsInfoArr := shopsInfo.SearchByUserId()

	goodsInfoArr := []model.GoodsModel{}
	for _, v := range shopsInfoArr {
		temp := v // 其实很简单 引入一个临时局部变量就可以了，这样就可以将每次的值存储到该变量地址上
		goodsInfo := model.GoodsModel{}
		initDB.Db.Model(&temp).Related(&goodsInfo, "GoodsId")
		goodsInfo.Num = temp.Nums
		goodsInfoArr = append(goodsInfoArr, goodsInfo)
	}

	result.Data = goodsInfoArr
	ctx.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

//根据商品类型获取商品

func SearchShopsByType(ctx *gin.Context) {
	shops := model.ShopsModel{}
	result := &model.ResultModel{
		Code:    200,
		Message: "查询成功",
		Data:    nil,
	}

	shopType := ctx.Query("type")

	shopInfoArr := shops.GetAllInfo()
	shopsInfoList := []model.ShopsModel{}

	for _, v := range shopInfoArr {
		temp := v // 其实很简单 引入一个临时局部变量就可以了，这样就可以将每次的值存储到该变量地址上
		shopsTemp := model.ShopsModel{}
		goods := model.GoodsModel{
			Id:  temp.GoodsId,
			Num: temp.Nums,
		}
		user := model.BussinessModel{
			Id: temp.BussinessId,
		}
		userInfo := user.GetOneBussinessInfo()
		userInfo.Password = "******"
		if goodsInfo := goods.FindById(); goodsInfo.Type == shopType {
			shopsTemp.Goods = goodsInfo
			shopsTemp.Bussiness = userInfo
			shopsInfoList = append(shopsInfoList, shopsTemp)
		}
	}
	result.Data = shopsInfoList
	ctx.JSON(http.StatusOK, gin.H{
		"result": result,
	})

}
