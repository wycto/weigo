package controller

import (
	"fmt"
	"io"
	"net/http"
)

type UserController struct {
	modelName string
}

func UserRouter(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "\nthis is UserRouter")
}

func (receiver *UserController) Login(w http.ResponseWriter, r *http.Request) string {
	return "this is User login"
}

func (receiver *UserController) Test() {
	fmt.Println("this is User login")
	r := http.Request{}
	fmt.Println("r:", r.URL.RawQuery)
	fmt.Println("r:", r.URL.Query())
	fmt.Println("r:", r.PostForm)
	fmt.Println("r:", r.FormValue("name"))
	fmt.Println("r:", r.Form.Get("name"))
}
