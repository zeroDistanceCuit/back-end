package model

type  GoodsInfoModel struct {
	Id int
	Model string
}

func (goodsInfo GoodsInfoModel) TableName() string {
	return "shops_info"
}
