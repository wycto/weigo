package controller

import (
	"io"
	"net/http"
)

type User struct {
	modelName string
}

func UserRouter(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "\nthis is UserRouter")
}

func (receiver User) Login(w http.ResponseWriter, r *http.Request) string {
	return "this is User login"
}
