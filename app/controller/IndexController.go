package controller

import (
	"wycto/weigo"
)

type IndexController struct {
	weigo.Controller
}

func (c *IndexController) Index() {
	c.Display("")
}
