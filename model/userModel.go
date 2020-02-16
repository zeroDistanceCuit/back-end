package model

import "back_end/config/initDB"

type UserModel struct {
	Id      int
	Name    string
	Phone   string
	Sex     string
	Address string
}

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

