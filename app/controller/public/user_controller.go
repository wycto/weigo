package public

import (
	"wycto/app/model"
	"wycto/weigo"
)

type UserController struct {
	weigo.Controller
}

func (c *UserController) Login() {
	if c.Context.IsPost() {
		c.ResponseSuccess("哈哈", c.Context.Request.Method)
	} else {
		c.ResponseError("不是post", c.Context.Request.Method)
	}
	return
}

func (c *UserController) Register() {
	if c.Context.IsPost() {
		model.UserModel().Select()
	} else {
		c.MethodNotAllowed()
	}
	return
}
