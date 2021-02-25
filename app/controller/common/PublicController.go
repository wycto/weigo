package common

import "wycto/weigo"

type PublicController struct {
	weigo.Controller
}

func (c *PublicController) Login() {
	if c.Context.IsPost() {
		c.Context.Success("哈哈", c.Context.Request.Method)
	} else {
		c.Context.Error("不是post", c.Context.Request.Method)
	}
	return
}

func (c *PublicController) Register() {

}
