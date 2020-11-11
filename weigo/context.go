package weigo

import "net/http"

type Context struct {
	Request        *http.Request
	ResponseWriter *Response
}
