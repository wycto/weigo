package controller

import (
	"fmt"
	"wycto/weigo"
)

type UserController struct {
	weigo.Controller
	modelName string
}

func (receiver *UserController) Login() {
	fmt.Println("this is user login", receiver.Context)
}

func (receiver *UserController) UserInfo() {
	fmt.Println("this is userinfo")
}
