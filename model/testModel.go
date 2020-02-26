package model

import (
	"back_end/config/initDB"
	"fmt"
)

type TestModel struct {
	Id      int  
	Message string
}

func (test TestModel) TableName() string {
	return "test"
}

func (test TestModel) Insert() int {
	create := initDB.Db.Create(&test)
	if create.Error != nil {
		return 0
	}
	return test.Id
}

func (test TestModel) FindAll()  []TestModel{
	var testArr []TestModel
	initDB.Db.Find(&testArr)
	return testArr
}

func (test TestModel)FindById() TestModel {
	initDB.Db.First(&test,test.Id)
	return  test
}

//更新所有字段
func (test TestModel)SaveHandler()  {
	fmt.Println(test)
	initDB.Db.Save(&test)
}

//更新单个字段

//fnc (data [string]string) Update()  {
////
////}u

func (test TestModel) Delete(){
	initDB.Db.Delete(test)
}

