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

func (receiver *User) Select() ([]map[string]interface{}, error) {
	rows, err := weigo.DataBase.Table("cto_controller").SetFields("name,id").Select()
	if err != nil {
		return nil, err
	}
	return rows, nil
}
