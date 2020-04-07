package model

import (
	"back_end/config/initDB"
)

type UserModel struct {
	Id      int
	UserModelId int
	Name    string
	Password string
	Phone   string
	Sex     string
	Address string
}

/**
* 判断数据库是否存在此表，无则生成
*/
//func init() {
//	if !initDB.Db.HasTable(UserModel{}) {
//		initDB.Db.CreateTable(UserModel{})
//	}
//}

func (user UserModel) TableName() string {
	return "user"
}

func (user UserModel) Insert() int {
	create:=initDB.Db.Create(&user)
	if create.Error!=nil{
		return 0
	}
	return user.Id
}

func (user UserModel) QueryByUsername() UserModel {
	var users UserModel
	initDB.Db.Where("name = ?", user.Name).Find(&users)
	return users
}

func (user UserModel)QueryById() UserModel {
	initDB.Db.First(&user,user.Id)
	return user
}