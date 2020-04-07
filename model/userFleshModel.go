package model

import (
	"back_end/config/initDB"
)

type UserFleshModel struct {
	Id int
	Model string
}

func (user UserFleshModel) TableName() string  {
	return "user_info"
}

func (user UserFleshModel) Find() UserFleshModel{
	initDB.Db.First(&user,user.Id)
	return user
}