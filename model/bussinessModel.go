package model

import (
	"back_end/config/initDB"
	"fmt"
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
	fmt.Println(user)
	initDB.Db.Where("name = ?", user.Name).Find(&users)
	fmt.Println("dasjdas",users)
	return users
}

