package model

import (
	"wycto/weigo"
)

func UserModel() *User {
	user := &User{}
	user.SetTableName("user")
	return user
}

type User struct {
	weigo.Model
}

func (receiver *User) Test() string {
	return "其他方法"
}
