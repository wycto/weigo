package weigo

import (
	"fmt"
	"net/http"
	"reflect"
)

type WeiContext struct {
	serveMux *http.Server
}

func Router(urlPath string, c interface{}) {
	serveMux := *http.NewServeMux()
	serveMux.HandleFunc("/", AppHandleFunc(c))
}

func AppHandleFunc(c interface{}) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		value := reflect.ValueOf(c)
		method := value.MethodByName("Index")
		if method.IsValid() {
			method.Call(nil)
		} else {
			fmt.Println("方法不存在")
		}
	}
}
