package route

import (
	"wycto/app/controller/index"
	"wycto/weigo"
)

func AddIndexRouter() {
	weigo.Router("/index/user/", &index.UserController{})
	weigo.Router("/index/api/", &index.APIController{})
	weigo.Router("/index/test/", &index.TestController{})
	weigo.Router("/index/model/", &index.ModelController{})
}
