package main

import (
	"fmt"
	"net/http"
)

func ttt(r http.Request) {
	fmt.Println("r.URL: ", r.URL)
	fmt.Println("r.URL.Path: ", r.URL.Path)
	fmt.Println("r.URL.Host: ", r.URL.Host)
	fmt.Println("r.URL.User: ", r.URL.User)
	fmt.Println("r.URL.ForceQuery: ", r.URL.ForceQuery)
	fmt.Println("r.URL.Fragment: ", r.URL.Fragment)
	fmt.Println("r.URL.Opaque: ", r.URL.Opaque)
	fmt.Println("r.URL.RawFragment: ", r.URL.RawFragment)
	fmt.Println("r.URL.RawPath: ", r.URL.RawPath)
	fmt.Println("r.URL.RawQuery: ", r.URL.RawQuery)
	fmt.Println("r.URL.Scheme: ", r.URL.Scheme)
	fmt.Println("r.URL.Query(): ", r.URL.Query())
	fmt.Println("r.URL.EscapedFragment(): ", r.URL.EscapedFragment())
	fmt.Println("r.URL.Hostname(): ", r.URL.Hostname())
	fmt.Println("r.URL.IsAbs(): ", r.URL.IsAbs())
	fmt.Println("r.URL.Port(): ", r.URL.Port())
	fmt.Println("r.URL.String(): ", r.URL.String())
	fmt.Println("-------------------------")
	fmt.Println("r.Host: ", r.Host)
	fmt.Println("r.Body: ", r.Body)
	fmt.Println("r.Method: ", r.Method)
	fmt.Println("r.Form: ", r.Form)
	fmt.Println("r.ContentLength: ", r.ContentLength)
	fmt.Println("r.Header: ", r.Header)
	fmt.Println("r.MultipartForm: ", r.MultipartForm)
	fmt.Println("r.PostForm: ", r.PostForm)
	fmt.Println("r.RemoteAddr: ", r.RemoteAddr)
	fmt.Println("r.RequestURI: ", r.RequestURI)
	fmt.Println("r.TransferEncoding: ", r.TransferEncoding)
	fmt.Println("r.Context(): ", r.Context())
	fmt.Println("r.FormValue(\"name\"): ", r.FormValue("name"))
	fmt.Println("r.ParseForm(): ", r.ParseForm())
	fmt.Println("----------------------------")

	fmt.Println("=============================")
	fmt.Println("   ")
	fmt.Println("   ")
}
