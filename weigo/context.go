package weigo

import "net/http"

type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	ControllerName string
	ActionName     string
	Param          map[string]string
}
