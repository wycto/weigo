package weigo

import (
	"fmt"
	"net/http"
)

/**
启动函数会执行一些初始化操作，如注初始化数据库、缓存、检测、路由注册等
*/
func Run() {

	//启动服务
	ip := GetIP()
	fmt.Println("服务已启动，地址：", ip+":9099")
	err := http.ListenAndServe(":9099", nil)
	if err != nil {
		fmt.Println("启动失败！！！")
	}
}
