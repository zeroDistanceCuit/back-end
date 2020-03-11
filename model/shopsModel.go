package model

import "back_end/config/initDB"

type ShopsModel struct {
	Id int
	BussinessId BussinessModel
	GoodsId int
	Nums int
}

func (shops ShopsModel) TableName() string {
	return "shops"
}

func (shops ShopsModel) Insert() int {
	create := initDB.Db.Create(&shops)
	if create.Error!=nil{
		return 0
	}
	return shops.Id
}

//联立查询
func (shops ShopsModel) FindById() ShopsModel {
	initDB.Db.First(&shops,shops.Id)
	return shops
}

func (shops ShopsModel) SearchByName() []ShopsModel {
	var shopsArr []ShopsModel
	//initDB.Db.Where("name=?",shops.)
	return shopsArr
}



