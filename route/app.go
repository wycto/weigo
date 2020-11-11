package route

import (
	"wycto/app/controller"
	"wycto/weigo"
)

func init() {
	weigo.Router("/", &controller.IndexController{})
}
