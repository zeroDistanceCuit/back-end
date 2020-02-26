package model

import (
	"back_end/config/initDB"
)

type BussinessModel struct {
	Id      int
	Name    string
	ShopName string
	Password string
}

func (user BussinessModel) TableName() string {
	return "bussiness"
}

func (user BussinessModel) Insert() int {
	create:=initDB.Db.Create(&user)
	if create.Error!=nil{
		return 0
	}
	return user.Id
}

func (user BussinessModel) QueryByUsername() BussinessModel {
	var users BussinessModel
	initDB.Db.Where("name = ?", user.Name).Find(&users)
	return users
}

//修改密码
func (user BussinessModel) UpdatePassW() bool {
	update:=initDB.Db.Model(&user).Update("password", user.Password)
	if update.Error !=nil{
		return  false
	}
	return true
}