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
	rows, err := weigo.GetDataBase().Table("cto_article").GetRows()
	if err != nil {
		fmt.Println()
	}
	json, _ := json.Marshal(rows)
	io.WriteString(c.Context.ResponseWriter, string(json))
}
