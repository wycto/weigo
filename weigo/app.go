package weigo

import (
	"io"
	"net/http"
	"reflect"
	"strings"
)

var WeiApp *App

type App struct {
	Server         *http.Server
	ControllerName string
	ActionName     string
}

func init() {
	WeiApp = NewApp()
}

func NewApp() *App {
	app := &App{Server: &http.Server{}}
	return app
}

func Router(urlPath string, c interface{}) {
	http.HandleFunc(urlPath, AppHandleFunc(c))
}

func AppHandleFunc(controller interface{}) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPath := r.URL.Path
		urlPathArr := strings.FieldsFunc(urlPath, func(r rune) bool {
			if r == '/' {
				return true
			} else {
				return false
			}
		})

		len := len(urlPathArr)
		if len > 1 {
			WeiApp.ActionName = urlPathArr[1]
		} else {
			WeiApp.ActionName = "Index"
		}

		reflectType := reflect.TypeOf(controller)
		value := reflect.ValueOf(controller)

		actionMap := make(map[string]string)
		for i := 0; i < reflectType.NumMethod(); i++ {
			m := reflectType.Method(i)
			n := m.Name
			actionMap[strings.ToLower(n)] = n
		}

		method := value.MethodByName(actionMap[strings.ToLower(WeiApp.ActionName)])
		if method.IsValid() {
			method.Call(nil)
		} else {
			io.WriteString(w, "请求地址错误"+r.URL.Path)
		}
	}
}
