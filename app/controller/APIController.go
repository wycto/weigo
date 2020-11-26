package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"wycto/weigo"
)

type APIController struct {
	weigo.Controller
}

func (c *APIController) Index() {
	var ww map[string]string
	ww = make(map[string]string)
	ww["uid"] = "2"
	ww["nickname"] = "[:string]"

	rows, err := weigo.DataBase.Name("user").Where(ww).GetAll()
	if err != nil {
		fmt.Println(err.Error())
	}
	json, _ := json.Marshal(rows)
	io.WriteString(c.Context.ResponseWriter, string(json))
}

func (c *APIController) Test() {
	re := RegexpWhereKey("age[NOT IN]")
	fmt.Println(re)
}

func RegexpWhereKey(Key string) string {
	reg, err := regexp.Compile("[>]|[<]|[=]|[<>]|[!=]|[LIKE]|[IN]|[^NOT IN$]")
	if err != nil {
		fmt.Println("regexp err:", err.Error())
		return ""
	}

	result := reg.FindAllStringSubmatch(Key, -1)
	fmt.Println(result)
	return ""
}
