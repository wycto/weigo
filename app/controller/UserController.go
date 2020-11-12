package controller

import (
	"fmt"
	"io"
	"wycto/weigo"
)

type UserController struct {
	weigo.Controller
	modelName string
}

func (c *UserController) Login() {
	fmt.Println("this is user login", c.Context.Request.Method)
	fmt.Println("name:", c.Context.Request.FormValue("name"))
	fmt.Println("name:", c.Context.Request.Form.Get("name"))
	fmt.Println("ControllerName:", c.Context.ControllerName)
	fmt.Println("ActionName:", c.Context.ActionName)
	io.WriteString(c.Context.ResponseWriter, "name:"+c.Context.Request.Form.Get("name"))
}

func (c *UserController) UserInfo() {
	fmt.Println("this is userinfo")
}
