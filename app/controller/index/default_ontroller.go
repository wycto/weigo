package index

import (
	"wycto/weigo"
)

type DefaultController struct {
	weigo.Controller
}

func (c *DefaultController) Index() {
	c.Assign("website", "http://www.wycto.cn")
	c.Assign("github", "http://www.github.com/wycto/weigo")
	c.Display("")
}
