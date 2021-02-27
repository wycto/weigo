package weigo

/*MVC的C层，控制器类*/
/*
控制器基类，框架控制器，业务控制器需要继承
*/
import (
	"html/template"
	"io"
	"net/http"
	"strings"
)

//控制器类
type Controller struct {
	Context *Context
	data    map[string]interface{}
}

//控制器初始化
func (controller *Controller) Init(context *Context) {
	controller.Context = context
	controller.data = make(map[string]interface{})
}

//控制器默认的请求方法
func (controller *Controller) Index() {
	http.NotFound(controller.Context.ResponseWriter, controller.Context.Request)
}

//列表方法
func (controller *Controller) List() {
	http.NotFound(controller.Context.ResponseWriter, controller.Context.Request)
}

//详情方法
func (controller *Controller) View() {
	http.NotFound(controller.Context.ResponseWriter, controller.Context.Request)
}

//删除方法
func (controller *Controller) Delete() {
	http.NotFound(controller.Context.ResponseWriter, controller.Context.Request)
}

//修改方法
func (controller *Controller) Update() {
	http.NotFound(controller.Context.ResponseWriter, controller.Context.Request)
}

//页面模版赋值
func (controller *Controller) Assign(Key string, Value interface{}) {
	controller.data[Key] = Value
}

//页面模版渲染
func (controller *Controller) Display(viewName string) {
	if viewName == "" {
		viewName = RootPath + DS + "view" + DS + strings.ToLower(controller.Context.AppName) + DS + strings.ToLower(controller.Context.ControllerName) + DS + strings.ToLower(controller.Context.ActionName) + ".html"
	}
	t, err := template.ParseFiles(viewName)
	if err != nil {
		io.WriteString(controller.Context.ResponseWriter, err.Error())
	} else {
		t.Execute(controller.Context.ResponseWriter, controller.data)
	}
}

//输出成功的json格式
func (controller *Controller) ResponseSuccess(msg string, data interface{}) {
	controller.Context.ResponseSuccess(msg, data)
}

//输出失败的json格式
func (controller *Controller) ResponseError(msg string, data interface{}) {
	controller.Context.ResponseError(msg, data)
}

//输出系统定义好的错误代码格式json数据
func (controller *Controller) ResponseErrorMessage(message *Message, data interface{}) {
	controller.Context.ResponseMessage(message, data)
}

//输出字符串
func (controller *Controller) ResponseString(msg string) {
	controller.Context.ResponseString(msg)
}

func (controller *Controller) MethodNotAllowed() {
	http.Error(controller.Context.ResponseWriter, "Method Not Allowed", http.StatusMethodNotAllowed)
}
