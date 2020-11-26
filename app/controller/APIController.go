package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"strings"
	"wycto/weigo"
)

type APIController struct {
	weigo.Controller
}

func (c *APIController) Index() {
	var ww map[string]string
	ww = make(map[string]string)
	ww["uid|<"] = "3"
	ww["nickname"] = "[:string]管理员"

	rows, err := weigo.DataBase.Name("user").Page(3, 2).GetAll()
	if err != nil {
		fmt.Println(err.Error())
	}
	json, _ := json.Marshal(rows)
	io.WriteString(c.Context.ResponseWriter, string(json))
}

func (c *APIController) Test() {
	Key, Reg := RegexpWhereKey("age|")
	fmt.Println(Key, Reg)
}

func RegexpWhereKey(Key string) (string, string) {
	reg, err := regexp.Compile(`\|(.*)`)
	if err != nil {
		fmt.Println("regexp err:", err.Error())
		return Key, "="
	}

	result := reg.FindAllString(Key, -1)
	if len(result) == 0 {
		return Key, "="
	}

	KeyIndexArr := reg.FindAllStringIndex(Key, -1)
	position := KeyIndexArr[0][0]
	field := Key[:position]
	regexpStr := Key[position+1:]
	regexpStr = strings.Trim(regexpStr, " ")
	if regexpStr == "" {
		regexpStr = "="
	}
	return field, regexpStr
}
