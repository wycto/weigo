package weigo

/*
此类为上下文，请求上下文
*/
import "net/http"

type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	AppName        string              //当前应用名称
	ControllerName string              //当前控制器名称
	ActionName     string              //当前动作方法名称
	paramData      map[string]string   //请求参数，合并类get、post
	getData        map[string]string   //get参数
	postData       map[string]string   //post参数
	Header         map[string][]string //header信息
}

//获取get参数
func (context *Context) Get(key string) string {
	val, err := context.getData[key]
	if err == false {
		return ""
	}
	return val
}

//获取post参数
func (context *Context) Post(key string) string {
	val, err := context.postData[key]
	if err == false {
		return ""
	}
	return val
}

//获取参数
func (context *Context) Param(key string) string {
	val, err := context.paramData[key]
	if err == false {
		return ""
	}
	return val
}

//所有get数据-数组
func (context *Context) GetData() map[string]string {
	return context.getData
}

//所有post 数据-数组
func (context *Context) PostData() map[string]string {
	return context.postData
}

//获取所有数据-数组
func (context *Context) ParamData() map[string]string {
	return context.paramData
}
