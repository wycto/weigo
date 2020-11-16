package controller

import (
	"wycto/weigo"
)

type IndexController struct {
	weigo.Controller
}

func (c *IndexController) Index() {
	c.Assign("name", "唯一")
	c.Display("")
}
