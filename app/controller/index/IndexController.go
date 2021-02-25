package index

import (
	"wycto/weigo"
)

type IndexController struct {
	weigo.Controller
}

func (c *IndexController) Index() {
	c.Assign("website", "http://www.wycto.cn")
	c.Assign("github", "http://www.github.com/wycto/go-frame")
	c.Display("")
}
