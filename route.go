package weigo

/*
路由类，解析路由和请求参数
*/
import (
	"encoding/json"
	"github.com/wycto/weigo/datatype"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path/filepath"
	"reflect"
	"strings"
)

//静态资源路由
func RouterStatic(url string, path string) {
	if path == "" {
		path = url
	}
	http.Handle(url, http.StripPrefix(url, http.FileServer(http.Dir(filepath.Join(RootPath, path)))))
}

//MVC控制器路由
func Router(urlPath string, c ControllerInterface) {
	http.HandleFunc(urlPath, AppHandleFunc(c))
}

func AppHandleFunc(controller ControllerInterface) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		var context = &Context{ResponseWriter: w, Request: r}
		context.Header = r.Header

		urlPath := r.URL.Path

		if urlPath == "/favicon.ico" {
			return
		}

		urlPathArr := strings.FieldsFunc(urlPath, func(r rune) bool {
			return r == '/'
		})

		get := datatype.Row{}

		length := len(urlPathArr)
		if length >= 3 {
			context.AppName = urlPathArr[0]
			context.ControllerName = urlPathArr[1]
			context.ActionName = urlPathArr[2]

			if length > 3 {
				//get参数处理
				for i := 3; i <= length; i = i + 2 {
					if i+1 < length {
						get[urlPathArr[i]] = urlPathArr[i+1]
					} else if i < length {
						get[urlPathArr[i]] = ""
					}
				}
			}
		} else if length == 2 {
			context.AppName = urlPathArr[0]
			context.ControllerName = urlPathArr[1]
			context.ActionName = Config.App.DefaultActionName
		} else if length == 1 {
			context.AppName = urlPathArr[0]
			context.ControllerName = Config.App.DefaultControllerName
			context.ActionName = Config.App.DefaultActionName
		} else {
			context.AppName = Config.App.DefaultAppName
			context.ControllerName = Config.App.DefaultControllerName
			context.ActionName = Config.App.DefaultActionName
		}

		//?后面的参数
		rawQueryMap, err := url.ParseQuery(r.URL.RawQuery)

		if err == nil {
			for i, v := range rawQueryMap {
				get[i] = v[len(v)-1]
			}
		}

		context.Get = &get

		//post参数处理
		jsonStr := ""
		ct := r.Header.Get("Content-Type")
		if len(ct) > 0 {
			if strings.Contains(ct, "multipart/form-data;") {
				r.ParseMultipartForm(32 << 20)
			} else if strings.Contains(ct, "application/x-www-form-urlencoded") {
				r.ParseForm()
			} else if strings.Contains(ct, "application/json") {

				con, err := ioutil.ReadAll(r.Body) //获取post的数据
				if err == nil {
					jsonStr = string(con)
				}
				r.Body.Close()
			}
		}

		//post数据
		post := datatype.Row{}

		form := r.PostForm
		for i, v := range form {
			post[i] = v[len(v)-1]
		}

		jsonData := make(map[string]string)
		json.Unmarshal([]byte(jsonStr), &jsonData)

		for k, val := range jsonData {
			post[k] = val
		}

		context.Post = &post

		//合并参数，post覆盖get
		param := datatype.Row{}
		for k, val := range *context.Get {
			param[k] = val
		}
		for k, val := range *context.Post {
			param[k] = val
		}
		context.Param = &param

		//反射调用
		reflectType := reflect.TypeOf(controller)
		value := reflect.ValueOf(controller)

		actionMap := make(map[string]string)
		for i := 0; i < reflectType.NumMethod(); i++ {
			m := reflectType.Method(i)
			n := m.Name
			actionMap[strings.ToLower(n)] = n
		}

		ActionName := actionMap[strings.ToLower(context.ActionName)]
		ControllerName := reflectType.Elem().Name()
		ControllerName = ControllerName[:len(ControllerName)-10]

		context.ControllerName = ControllerName
		context.ActionName = ActionName
		controller.Init(context)

		method := value.MethodByName(ActionName)

		if method.IsValid() {
			method.Call(nil)
		} else {
			io.WriteString(w, "404 page not found ："+r.URL.Path)
		}
	}
}
