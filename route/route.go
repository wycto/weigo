package route

import (
	"wycto/app/controller/index"
	"wycto/weigo"
)

func init() {
	//路由定义必须遵循MVC：/public/public/  代表common应用（模块）、public控制器
	weigo.Router("/", &index.DefaultController{})
	AddPublicRouter()
	AddIndexRouter() //前台应用路由定义
	AddAdminRouter() //后台应用路由定义
}
