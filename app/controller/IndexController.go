package controller

import (
	"fmt"
	"wycto/weigo"
)

type IndexController struct {
	weigo.Controller
}

func (c *IndexController) Index() {
	fmt.Println("Welcome WeiGo；c:" + c.Context.ControllerName + ";a:" + c.Context.ActionName)
}
