package controller

import (
	"html/template"
	"io"
	"wycto/weigo"
)

type UserController struct {
	weigo.Controller
	modelName string
}

func (c UserController) Index() {
	io.WriteString(c.Context.ResponseWriter, "Welcome UserController Index")
}

func (c *UserController) Login() {
	io.WriteString(c.Context.ResponseWriter, "this is user login"+c.Context.Request.Method)
	io.WriteString(c.Context.ResponseWriter, "name:"+c.Context.Request.FormValue("name"))
	io.WriteString(c.Context.ResponseWriter, "name:"+c.Context.Request.Form.Get("name"))
	io.WriteString(c.Context.ResponseWriter, "ControllerName:"+c.Context.ControllerName)
	io.WriteString(c.Context.ResponseWriter, "ActionName:"+c.Context.ActionName)
	io.WriteString(c.Context.ResponseWriter, "name:"+c.Context.Request.Form.Get("name"))
}

func (c *UserController) UserInfo() {
	t, err := template.ParseFiles("app/view/test.html")
	if err != nil {
		io.WriteString(c.Context.ResponseWriter, err.Error())
	} else {
		t.Execute(c.Context.ResponseWriter, c)
	}

	io.WriteString(c.Context.ResponseWriter, "\nMethod:"+c.Context.Request.Method)
	io.WriteString(c.Context.ResponseWriter, "\nname:"+c.Context.GetParam("name"))
	io.WriteString(c.Context.ResponseWriter, "\n Request.PostForm.Get(\"name2\"):"+c.Context.Request.PostForm.Get("name2"))
	io.WriteString(c.Context.ResponseWriter, "\n Request.PostFormValue(\"name3\"):"+c.Context.Request.PostFormValue("name3"))
	io.WriteString(c.Context.ResponseWriter, "\n Request.Form.Get(\"name4\"):"+c.Context.Request.Form.Get("name4"))
	io.WriteString(c.Context.ResponseWriter, "\n Request.Form.Encode():"+c.Context.Request.Form.Encode())
	io.WriteString(c.Context.ResponseWriter, "\n Request.URL.RawQuery:"+c.Context.Request.URL.RawQuery)
	io.WriteString(c.Context.ResponseWriter, "\n Request.URL.Query().Encode():"+c.Context.Request.URL.Query().Encode())

}
