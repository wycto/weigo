package weigo

import (
	"net/http"
)

var WeiApp *App
var WeiContext *Context

type App struct {
	Server         *http.Server
	ControllerName string
	ActionName     string
}

func init() {
	WeiApp = NewApp()
}

func NewApp() *App {
	app := &App{Server: &http.Server{}}
	return app
}
