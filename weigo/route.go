package weigo

import (
	"io"
	"net/http"
	"reflect"
	"strings"
)

func Router(urlPath string, c ControllerInterface) {
	http.HandleFunc(urlPath, AppHandleFunc(c))
}

func AppHandleFunc(controller ControllerInterface) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		Context := &Context{ResponseWriter: w, Request: r}

		urlPath := r.URL.Path
		urlPathArr := strings.FieldsFunc(urlPath, func(r rune) bool {
			if r == '/' {
				return true
			} else {
				return false
			}
		})

		length := len(urlPathArr)
		if length > 1 {
			Context.ActionName = urlPathArr[1]
			if length > 2 {
				param := make(map[string]string)
				for i := 2; i <= length; i = i + 2 {
					if i+1 < length {
						param[urlPathArr[i]] = urlPathArr[i+1]
					} else if i < length {
						param[urlPathArr[i]] = ""
					}
				}

				Context.Param = param
				Context.getData = param
			}
		} else {
			Context.ActionName = "Index"
		}

		reflectType := reflect.TypeOf(controller)
		value := reflect.ValueOf(controller)

		actionMap := make(map[string]string)
		for i := 0; i < reflectType.NumMethod(); i++ {
			m := reflectType.Method(i)
			n := m.Name
			actionMap[strings.ToLower(n)] = n
		}

		ActionName := actionMap[strings.ToLower(Context.ActionName)]
		ControllerName := reflectType.Elem().Name()
		ControllerName = ControllerName[:len(ControllerName)-10]

		Context.ControllerName = ControllerName
		Context.ActionName = ActionName
		controller.Init(Context)

		method := value.MethodByName(ActionName)

		if method.IsValid() {
			method.Call(nil)
		} else {
			io.WriteString(w, "404 page not found ："+r.URL.Path)
		}
	}
}
