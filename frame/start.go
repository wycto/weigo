package frame

import (
	"fmt"
	"io"
	"net/http"
)

/**
启动函数会执行一些初始化操作，如注初始化数据库、缓存、检测、路由注册等
*/
func Run() {
	//注册路由
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", Index)
	//启动服务
	fmt.Println("启动 端口为：9099 的服务")
	err := http.ListenAndServe(":9099", serveMux)
	if err != nil {
		fmt.Println("启动服务器失败")
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "这是首页")
	io.WriteString(w, r.Method)
	io.WriteString(w, r.Host)
	fmt.Println("Method: ", r.Method)
	fmt.Println("URL: ", r.URL)
	fmt.Println("header: ", r.Header)
	fmt.Println("body: ", r.Body)
	fmt.Println("RemoteAddr: ", r.RemoteAddr)
}
