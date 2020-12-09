package weigo

import "net/http"

type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	ControllerName string
	ActionName     string
	paramData      map[string]string
	getData        map[string]string
	postData       map[string]string
	Header         map[string][]string
}

func (context *Context) Get(key string) string {
	val, err := context.getData[key]
	if err == false {
		return ""
	}
	return val
}

func (context *Context) Post(key string) string {
	val, err := context.postData[key]
	if err == false {
		return ""
	}
	return val
}

func (context *Context) Param(key string) string {
	val, err := context.paramData[key]
	if err == false {
		return ""
	}
	return val
}

/**
get数据
*/
func (context *Context) GetData() map[string]string {
	return context.getData
}

/**
post 数据
*/
func (context *Context) PostData() map[string]string {
	return context.postData
}

/**
所有数据
*/
func (context *Context) ParamData() map[string]string {
	return context.paramData
}
