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

	rows, errorStr := weigo.DataBase.Name("user").SetFields("email,`name`,`nickname`").Where(ww).GetAll()
	if errorStr != "" {
		fmt.Println(errorStr)
	}
	json, _ := json.Marshal(rows)
	io.WriteString(c.Context.ResponseWriter, string(json))
}

func (c *APIController) DeleteAll() {
	var ww map[string]interface{}
	ww = make(map[string]interface{})
	ww["uid|<"] = 5
	ww["nickname"] = "唯一"

	rows, errorStr := weigo.DataBase.Name("user").DeleteAll()
	if errorStr != "" {
		fmt.Println(errorStr)
	}
	json, _ := json.Marshal(rows)
	io.WriteString(c.Context.ResponseWriter, string(json))
}

func (c *APIController) Delete() {
	var ww map[string]interface{}
	ww = make(map[string]interface{})
	ww["nickname"] = "update后的唯一"

	rows, errorStr := weigo.DataBase.Name("user").Where(ww).Delete()
	if errorStr != "" {
		fmt.Println(errorStr)
	}
	json, _ := json.Marshal(rows)
	io.WriteString(c.Context.ResponseWriter, string(json))
}

func (c *APIController) Update() {
	var ww map[string]interface{}
	ww = make(map[string]interface{})
	ww["nickname"] = "唯一"

	var dd map[string]interface{}
	dd = make(map[string]interface{})
	dd["nickname"] = "update后的唯一"

	rows, errorStr := weigo.DataBase.Name("user").Where(ww).Update(dd)
	if errorStr != "" {
		fmt.Println(errorStr)
	}
	json, _ := json.Marshal(rows)
	io.WriteString(c.Context.ResponseWriter, string(json))
}

func (c *APIController) Add() {
	var ww map[string]interface{}
	ww = make(map[string]interface{})
	ww["uid"] = 3

	var dd map[string]interface{}
	dd = make(map[string]interface{})
	dd["nickname"] = "唯一"

	rows, errorStr := weigo.DataBase.Name("user").Insert(dd)
	if errorStr != "" {
		fmt.Println(errorStr)
	}
	json, _ := json.Marshal(rows)
	io.WriteString(c.Context.ResponseWriter, string(json))
}

func (c *APIController) UpdateAll() {
	var ww map[string]interface{}
	ww = make(map[string]interface{})
	ww["uid"] = 3

	var dd map[string]interface{}
	dd = make(map[string]interface{})
	dd["nickname"] = "唯一333"

	rows, errorStr := weigo.DataBase.Name("user").Where(ww).UpdateAll(dd)
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
