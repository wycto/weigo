package route

import (
	"wycto/app/controller/common"
	"wycto/app/controller/index"
	"wycto/weigo"
)

func init() {
	//路由定义必须遵循MVC：/common/public/  代表common应用（模块）、public控制器
	weigo.Router("/", &index.IndexController{})
	weigo.Router("/common/public/", &common.PublicController{})
	AddIndexRouter() //前台应用路由定义
	AddAdminRouter() //后台应用路由定义
}
