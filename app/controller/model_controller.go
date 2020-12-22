package controller

import (
	"fmt"
	"wycto/app/model"
	"wycto/weigo"
)

type ModelController struct {
	weigo.Controller
}

func (receiver *ModelController) GetOne() {
	row, _ := model.UserModel().Find()
	fmt.Println(row)

	fmt.Println(model.UserModel().Test())
}
