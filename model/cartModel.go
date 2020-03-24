package model

import (
	"back_end/config/initDB"
	_ "golang.org/x/tools/go/ssa/interp"
)

type CartModel struct {
	Id int
	UserId int
	User UserModel
	BussinessId int
	Bussiness BussinessModel
	Num int
	Status string
	ShopsId int
	Shops ShopsModel
	Order string
	Time string
	//time.Time

}

func (cart CartModel)TableName() string {
	return "cart"
}

func (cart CartModel) Insert() int {
	create:=initDB.Db.Create(&cart)
	if create.Error!=nil{
		return 0
	}
	return cart.Id
}

func (cart CartModel)SearchByUserId() []CartModel {
	var cartList []CartModel
	initDB.Db.Where("user_id=? and status=?",cart.UserId,cart.Status).Find(&cartList)
	return cartList
}

//付款需要的信息
func (cart CartModel)UpdateByOrder() bool {
	update:=initDB.Db.Model(&cart).Update(CartModel{Status:cart.Status,Order:cart.Order,Time:cart.Time,Num:cart.Num})
	if update.Error!=nil{
		return false
	}else {
		return true
	}

}

func (cart CartModel)Delete() bool{
	if cart.Id !=0 {
		delete:=initDB.Db.Delete(&cart)
		if delete.Error!=nil{
			return false
		}else {
			return true
		}
	}else{
		return false
	}
}
