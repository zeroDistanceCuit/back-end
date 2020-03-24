package model

import (
	"back_end/config/initDB"
	"fmt"
)

type ShopsModel struct {
	Id          int
	BussinessId int
	Bussiness   BussinessModel
	GoodsId     int
	Goods       GoodsModel
	Nums        int
}

func (shops ShopsModel) TableName() string {
	return "shops"
}

func (shops ShopsModel) Insert() int {
	create := initDB.Db.Create(&shops)
	if create.Error != nil {
		return 0
	}
	return shops.Id
}

func (shops ShopsModel) Search() ShopsModel {
	var shopsEx ShopsModel
	initDB.Db.Where("bussiness_id=? and goods_id=?", shops.BussinessId, shops.GoodsId).Find(&shopsEx)
	if shopsEx.BussinessId > 0 {
		return shopsEx
	} else {
		id := shops.Insert()
		shops.Id = id
		shops.Nums = 0
		return shops
	}
}

// 根据用户id查询
func (shops ShopsModel) SearchByUserId() []ShopsModel {
	var shopsEx []ShopsModel
	initDB.Db.Where("bussiness_id=? ", shops.BussinessId).Find(&shopsEx)
	return shopsEx
}

//更新商品数量
func (shops ShopsModel) Update() bool {
	update := initDB.Db.Model(&shops).Update("nums", shops.Nums)
	if update.Error != nil {
		return false
	}
	return true
}

//获得所有信息
func (shops ShopsModel) GetAllInfo() []ShopsModel {
	var shopInfoAll []ShopsModel
	initDB.Db.Find(&shopInfoAll)
	return shopInfoAll
}

//根据Id查询
func (shops ShopsModel) GetById() ShopsModel {
	fmt.Println(shops)
	initDB.Db.First(&shops,shops.Id)
	return shops
}