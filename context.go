package weigo

/*
此类为上下文，请求上下文
*/
import (
	"encoding/json"
	"github.com/wycto/weigo/datatype"
	"net/http"
)

type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	AppName        string              //当前应用名称
	ControllerName string              //当前控制器名称
	ActionName     string              //当前动作方法名称
	Param          *datatype.Row       //请求参数，合并类get、post
	Get            *datatype.Row       //get参数
	Post           *datatype.Row       //post参数
	Header         map[string][]string //header信息
}

//是post请求
func (context *Context) IsGet() bool {
	if context.Request.Method == "GET" {
		return true
	}
	return false
}

//是post请求
func (context *Context) IsPost() bool {
	if context.Request.Method == "POST" {
		return true
	}
	return false
}

//参数存在返回true
func (context *Context) Has(keys ...string) bool {

	for _, key := range keys {
		_, ok := (*context.Param)[key]
		if !ok {
			return false
		}
	}

	return true
}

//参数存在且不为空返回true
func (context *Context) HasAndNotEmpty(keys ...string) bool {

	for _, key := range keys {
		v, ok := (*context.Param)[key]
		if !ok {
			return false
		} else if v == "" || v == "false" || v == "null" || v == "0" {
			return false
		}
	}

	return true
}

//参数不存在返回true
func (context *Context) NotHas(keys ...string) bool {
	b := context.Has(keys...)
	return !b
}

//不存在或者为空返回true
func (context *Context) NotHasOrEmpty(keys ...string) bool {
	b := context.HasAndNotEmpty(keys...)
	return !b
}

//响应字符串
func (context *Context) ResponseString(str string) {
	context.ResponseWriter.Write([]byte(str))
}

//响应json数据
func (context *Context) ResponseJson(data interface{}) {
	json, err := json.Marshal(data)
	if err != nil {
		context.ResponseWriter.Write([]byte(""))
	} else {
		context.ResponseWriter.Write(json)
	}
}

//响应code、msg json
func (context *Context) ResponseApiJson(code int, msg string, data interface{}) {

	dataMap := make(map[string]interface{})
	dataMap["code"] = code
	dataMap["msg"] = msg
	dataMap["data"] = data

	context.ResponseJson(dataMap)
}

//输出错误json
func (context *Context) ResponseError(msg string, data interface{}) {
	context.ResponseApiJson(1, msg, data)
}

//输出成功json
func (context *Context) ResponseSuccess(msg string, data interface{}) {
	context.ResponseApiJson(0, msg, data)
}

//响应系统级信息
func (context *Context) ResponseMessage(message *Message, data interface{}) {
	context.ResponseApiJson(message.Code, message.Msg, data)
}
