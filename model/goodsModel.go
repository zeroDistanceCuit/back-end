package model

import (
	"back_end/config/initDB"
	"fmt"
	_ "reflect"
	"regexp"
)

type GoodsModel struct {
	Id           int
	GoodsModelId int
	Name         string
	Type         string
	Money        string
	Num          int
	Img 		string
}

func (goods GoodsModel) TableName() string {
	return "goods"
}

// TODO 可以加入一个关键字组进行更精准过滤
func (goods GoodsModel) FindByName() []GoodsModel {
	var goodsArr []GoodsModel
	var result []GoodsModel
	initDB.Db.Where("type=?", goods.Type).Find(&goodsArr)

	if goods.Name == ""{
		for _,v:=range goodsArr{
			temp:=v
			result=append(result,temp)
		}
	}else{
	reg, _ := regexp.Compile(goods.Name)
	for _,v:=range goodsArr{
		temp:=v
		resFlag:=reg.FindAllString(temp.Name,-1)

		if len(resFlag) != 0{
			result=append(result,temp)
		}
	}
	}
	fmt.Println(result)
	return result
}

//仅仅关键字进行查询
func (goods GoodsModel) FindOnlyByName() []GoodsModel {
	var goodsArr []GoodsModel
	var result []GoodsModel
	reg, _ := regexp.Compile(goods.Name)
	initDB.Db.Find(&goodsArr)
	for _,v:=range goodsArr{
		temp:=v
		resFlag:=reg.FindAllString(temp.Name,-1)

		if len(resFlag) != 0{
			result=append(result,temp)
		}
	}

	return result
}

func (goods GoodsModel) FindById() GoodsModel {
	initDB.Db.First(&goods, goods.Id)
	return goods
}
