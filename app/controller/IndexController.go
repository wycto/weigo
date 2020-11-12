package controller

import (
	"io"
	"wycto/weigo"
)

type IndexController struct {
	weigo.Controller
}

func (c *IndexController) Index() {
	io.WriteString(c.Context.ResponseWriter, "Welcome WeiGo；c:"+c.Context.ControllerName+";a:"+c.Context.ActionName)
}
