package controller

import (
	"wycto/app/model"
	"wycto/weigo"
)

type IndexController struct {
	weigo.Controller
}

func (c *IndexController) Index() {
	user := model.User{}
	user.Select()
	c.Assign("name", "唯一")
	c.Display("")
}
