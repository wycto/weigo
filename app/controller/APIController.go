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
	var ww map[string]interface{}
	ww = make(map[string]interface{})
	ww["uid|<"] = 5
	ww["nickname"] = c

	rows, errorStr := weigo.DataBase.Name("user").SetFields("email,`name`,`nickname`").Where(ww).GetAll()
	if errorStr != "" {
		fmt.Println(errorStr)
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
