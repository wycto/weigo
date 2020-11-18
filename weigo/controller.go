package weigo

import (
	"html/template"
	"io"
	"net/http"
	"strings"
)

type Controller struct {
	Context *Context
	data    map[string]interface{}
}

func (controller *Controller) Init(context *Context) {
	controller.Context = context
	controller.data = make(map[string]interface{})
}

func (controller *Controller) Index() {
	http.Error(controller.Context.ResponseWriter, "Method Not Allowed", http.StatusMethodNotAllowed)
}

func (controller *Controller) List() {
	http.Error(controller.Context.ResponseWriter, "Method Not Allowed", http.StatusMethodNotAllowed)
}

func (controller *Controller) View() {
	http.Error(controller.Context.ResponseWriter, "Method Not Allowed", http.StatusMethodNotAllowed)
}

func (controller *Controller) Delete() {
	http.Error(controller.Context.ResponseWriter, "Method Not Allowed", http.StatusMethodNotAllowed)
}

func (controller *Controller) Update() {
	http.Error(controller.Context.ResponseWriter, "Method Not Allowed", http.StatusMethodNotAllowed)
}

func (controller *Controller) Assign(Key string, Value interface{}) {
	controller.data[Key] = Value
}
func (controller *Controller) Display(viewName string) {
	if viewName == "" {
		viewName = "app/view/" + strings.ToLower(controller.Context.ControllerName) + "/" + strings.ToLower(controller.Context.ActionName) + ".html"
	}
	t, err := template.ParseFiles(viewName)
	if err != nil {
		io.WriteString(controller.Context.ResponseWriter, err.Error())
	} else {
		t.Execute(controller.Context.ResponseWriter, controller.data)
	}
}
