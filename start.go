package weigo

/*
系统启用入口
*/
import (
	"fmt"
	"net/http"
)

/**
启动函数会执行一些初始化操作，如注初始化数据库、缓存、检测、路由注册等
路由在路由注册文件里面，不在这儿，移步到route文件夹下的route.go文件
*/
func Run() {

	//初始化配置
	ConfigInit("")

	//启动服务
	ip := GetIP()
	fmt.Println("服务已启动，地址：", ip+":"+Config.App.ServerPort+" OR 127.0.0.1:"+Config.App.ServerPort)
	err := http.ListenAndServe(":"+Config.App.ServerPort, nil)
	fmt.Println(err)
	if err != nil {
		fmt.Println("启动失败！！！")
	}
}
