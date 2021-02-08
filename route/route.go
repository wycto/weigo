package route

import (
	"wycto/app/index/controller"
	"wycto/weigo"
)

func init() {
	weigo.Router("/", &controller.IndexController{})
	weigo.Router("/index/user/", &controller.UserController{})
	weigo.Router("/index/api/", &controller.APIController{})
	weigo.Router("/index/test/", &controller.TestController{})
	weigo.Router("/index/model/", &controller.ModelController{})
}
