package model

import (
	"wycto/weigo"
)

func UserModel() *IndexModelUser {
	user := &IndexModelUser{}
	user.SetTableName("user")
	return user
}

type IndexModelUser struct {
	weigo.Model
}

func (receiver *IndexModelUser) Test() string {
	return "其他方法"
}
