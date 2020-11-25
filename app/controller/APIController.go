package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"wycto/weigo"
)

type APIController struct {
	weigo.Controller
}

func (c *APIController) Index() {
	var ww map[string]string
	ww = make(map[string]string)
	ww["uid"] = "2"
	ww["nickname"] = "超级管理员"

	rows, err := weigo.DataBase.Name("user").Where(ww).GetAll()
	if err != nil {
		fmt.Println(err.Error())
	}
	json, _ := json.Marshal(rows)
	io.WriteString(c.Context.ResponseWriter, string(json))
}

func (c *APIController) Test() {
	var where map[string]string
	where = make(map[string]string)
	where["name"] = "weiyi"
	BBB(where)

	var a = 10
	BBB(a)

	var b = "10"
	BBB(b)

	var d = 10.25
	BBB(d)

	var e = "weiyi"
	BBB(e)

	var f = [5]int{23, 67, 89, 23, 34}
	BBB(f)

	s := []int{1, 2, 3}
	BBB(s)
}

func BBB(Where interface{}) {
	ValueOf := reflect.ValueOf(Where)
	fmt.Println(ValueOf.Kind().String())
	fmt.Println(ValueOf.String())
	fmt.Println("--------")

}
