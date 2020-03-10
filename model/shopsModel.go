package model

type shopsModel struct {
	Id int
	BussinessId int `gorm:"index"`
	GoodsId int
	Nums int
}

func (shops shopsModel) TableName() string {
	return "shops"
}

