package model

import "back_end/config/initDB"

type  GoodsInfoModel struct {
	Id int
	Model string
}

func (goodsInfo GoodsInfoModel) TableName() string {
	return "goods_info"
}

func (good GoodsInfoModel) Find() GoodsInfoModel{
	initDB.Db.First(&good,good.Id)
	return good
}
