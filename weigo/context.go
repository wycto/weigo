package weigo

import "net/http"

type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	ControllerName string
	ActionName     string
	Param          map[string]string
	getData        map[string]string
	postData       map[string]string
}

func (context *Context) GetParam(key string) string {
	val, err := context.Param[key]
	if err == false {
		return ""
	}
	return val
}

func (context *Context) SetParam(key string, val string) {
	context.Param["key"] = "val"
}
