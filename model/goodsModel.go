package model

import (
	"back_end/config/initDB"
)

type GoodsModel struct {
	Id           int
	GoodsModelId int
	Name         string
	Type         string
	Money        string
	Num          int
}

func (goods GoodsModel) TableName() string {
	return "goods"
}

func (goods GoodsModel) FindByName() []GoodsModel {
	var goodsArr []GoodsModel
	initDB.Db.Where("name=? and type=?", goods.Name, goods.Type).Find(&goodsArr)
	return goodsArr
}

func (goods GoodsModel) FindById() GoodsModel {
	initDB.Db.First(&goods, goods.Id)
	return goods
}
