package weigo

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

func Router(urlPath string, c ControllerInterface) {
	http.HandleFunc(urlPath, AppHandleFunc(c))
}

func AppHandleFunc(controller ControllerInterface) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		Context := &Context{ResponseWriter: w, Request: r}
		Context.Header = r.Header

		urlPath := r.URL.Path
		urlPathArr := strings.FieldsFunc(urlPath, func(r rune) bool {
			return r == '/'
		})

		get := make(map[string]string)

		length := len(urlPathArr)
		if length > 1 {
			Context.ActionName = urlPathArr[1]
			if length > 2 {
				//get参数处理
				for i := 2; i <= length; i = i + 2 {
					if i+1 < length {
						get[urlPathArr[i]] = urlPathArr[i+1]
					} else if i < length {
						get[urlPathArr[i]] = ""
					}
				}

				//?后面的参数
				rawQueryMap, err := url.ParseQuery(r.URL.RawQuery)

				if err == nil {
					for i, v := range rawQueryMap {
						get[i] = v[len(v)-1]
					}
				}

				Context.getData = get
			}
		} else {
			Context.ActionName = "Index"
		}

		//post参数处理
		jsonStr := ""
		ct := r.Header.Values("Content-Type")
		if len(ct) > 0 {
			if strings.Contains(ct[0], "multipart/form-data;") {
				r.ParseMultipartForm(32 << 20)
			} else if ct[0] == "application/x-www-form-urlencoded" {
				r.ParseForm()
			} else if ct[0] == "application/json" {

				con, err := ioutil.ReadAll(r.Body) //获取post的数据
				if err == nil {
					jsonStr = string(con)
				}
				r.Body.Close()
			}
		}

		post := make(map[string]string)

		form := r.PostForm
		for i, v := range form {
			post[i] = v[len(v)-1]
		}

		jsonData := make(map[string]string)
		json.Unmarshal([]byte(jsonStr), &jsonData)

		for k, val := range jsonData {
			post[k] = val
		}

		Context.postData = post

		//合并参数，post覆盖get
		param := make(map[string]string)
		for k, val := range Context.getData {
			param[k] = val
		}
		for k, val := range Context.postData {
			param[k] = val
		}
		Context.paramData = param

		//反射调用
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
