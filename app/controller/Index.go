package controller

import (
	"io"
	"net/http"
)

type Index struct {
	modelName string
}

func IndexController(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "\n r.URL.String(): "+r.URL.String())
	io.WriteString(w, "\n r.Method: "+r.Method)
	io.WriteString(w, "\n r.Host: "+r.Host)
	io.WriteString(w, "\n r.Proto: "+r.Proto)
	io.WriteString(w, "\n r.RemoteAddr: "+r.RemoteAddr)
	io.WriteString(w, "\n r.RequestURI: "+r.RequestURI)
	io.WriteString(w, "\n r.URL.Host: "+r.URL.Host)
	io.WriteString(w, "\n r.URL.Fragment: "+r.URL.Fragment)
	io.WriteString(w, "\n r.URL.Path: "+r.URL.Path)
	io.WriteString(w, "\n r.URL.RawPath: "+r.URL.RawPath)
	io.WriteString(w, "\n r.URL.RawQuery: "+r.URL.RawQuery)
	io.WriteString(w, "\n r.URL.Hostname(): "+r.URL.Hostname())
	io.WriteString(w, "\n r.URL.Port(): "+r.URL.Port())
	io.WriteString(w, "\n r.Form.Get(\"aa\"): "+r.Form.Get("aa"))
	io.WriteString(w, "\n r.FormValue(\"aa\"): "+r.FormValue("aa"))
}

func (receiver *Index) Index() {

}
