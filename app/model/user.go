package model

import (
	"wycto/weigo"
)

type User struct {
	Name     string
	Age      int
	Like     []string
	UserInfo map[string]interface{}
}

func (receiver *User) Select() ([]map[string]interface{}, string) {
	rows, errorStr := weigo.DataBase.Table("cto_controller").SetFields("name,id").GetAll()
	if errorStr != "" {
		return nil, errorStr
	}
	return rows, ""
}
