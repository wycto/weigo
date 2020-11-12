package weigo

import (
	"net/http"
)

type Controller struct {
	Context *Context
}

func (controller *Controller) Init(context *Context) {
	controller.Context = context
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
