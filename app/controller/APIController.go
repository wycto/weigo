package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"wycto/weigo"
)

type APIController struct {
	weigo.Controller
}

func (c *APIController) Index() {
	rows, err := weigo.DataBase.Name("user").Where("uid<5").GetAll()
	if err != nil {
		fmt.Println(err.Error())
	}
	json, _ := json.Marshal(rows)
	io.WriteString(c.Context.ResponseWriter, string(json))
}

func (c *APIController) Test() {

	io.WriteString(c.Context.ResponseWriter, "Database:"+weigo.Config.DB.Database)
	io.WriteString(c.Context.ResponseWriter, "UserName:"+weigo.Config.DB.UserName)
}
