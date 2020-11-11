package weigo

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

/**
启动函数会执行一些初始化操作，如注初始化数据库、缓存、检测、路由注册等
*/
func Run() {

	//启动服务
	fmt.Println("启动 端口为：9099 的服务")
	err := http.ListenAndServe(":9099", nil)
	if err != nil {
		fmt.Println("启动失败！！！")
	}
}

func route(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "路由进来的")
	io.WriteString(w, "\n r.RequestURI: "+r.RequestURI)
	io.WriteString(w, "\n r.URL.String(): "+r.URL.String())
	io.WriteString(w, "\n r.Method: "+r.Method)
	io.WriteString(w, "\n r.Host: "+r.Host)
	io.WriteString(w, "\n r.Proto: "+r.Proto)
	io.WriteString(w, "\n r.RemoteAddr: "+r.RemoteAddr)
	io.WriteString(w, "\n r.RequestURI: "+r.RequestURI)
	io.WriteString(w, "\n r.URL.Host: "+r.URL.Host)
	io.WriteString(w, "\n r.URL.Fragment: "+r.URL.Fragment)
	io.WriteString(w, "\n r.URL.Path: "+r.URL.Path)
	io.WriteString(w, "\n r.URL.RawPath: "+r.URL.RawPath)
	io.WriteString(w, "\n r.URL.RawQuery: "+r.URL.RawQuery)
	io.WriteString(w, "\n r.URL.Hostname(): "+r.URL.Hostname())
	io.WriteString(w, "\n r.URL.Port(): "+r.URL.Port())
	io.WriteString(w, "\n r.Form.Get(\"aa\"): "+r.Form.Get("aa"))
	io.WriteString(w, "\n r.FormValue(\"aa\"): "+r.FormValue("aa"))

	s := r.URL.Path
	ca := strings.FieldsFunc(s, func(r rune) bool {
		if r == '/' {
			return true
		} else {
			return false
		}
	})

	fmt.Printf("\n%v", ca)
	fmt.Printf("\n%v", r)
}
