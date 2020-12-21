package weigo

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
	http.Error(controller.Context.ResponseWriter, "Method Not Allowed", http.StatusMethodNotAllowed)
}

//列表方法
func (controller *Controller) List() {
	http.Error(controller.Context.ResponseWriter, "Method Not Allowed", http.StatusMethodNotAllowed)
}

//详情方法
func (controller *Controller) View() {
	http.Error(controller.Context.ResponseWriter, "Method Not Allowed", http.StatusMethodNotAllowed)
}

//删除方法
func (controller *Controller) Delete() {
	http.Error(controller.Context.ResponseWriter, "Method Not Allowed", http.StatusMethodNotAllowed)
}

//修改方法
func (controller *Controller) Update() {
	http.Error(controller.Context.ResponseWriter, "Method Not Allowed", http.StatusMethodNotAllowed)
}

//页面模版赋值
func (controller *Controller) Assign(Key string, Value interface{}) {
	controller.data[Key] = Value
}

//页面模版渲染
func (controller *Controller) Display(viewName string) {
	if viewName == "" {
		viewName = RootPath + Config.View.RootPath + "/" + strings.ToLower(controller.Context.ControllerName) + "/" + strings.ToLower(controller.Context.ActionName) + ".html"
	}
	t, err := template.ParseFiles(viewName)
	if err != nil {
		io.WriteString(controller.Context.ResponseWriter, err.Error())
	} else {
		t.Execute(controller.Context.ResponseWriter, controller.data)
	}
}
