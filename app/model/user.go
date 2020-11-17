package model

import (
	"fmt"
	"wycto/weigo"
)

type User struct {
	Name     string
	Age      int
	Like     []string
	UserInfo map[string]interface{}
}

func (receiver *User) Select() {
	rows, err := weigo.GetDataBase().DB.Query("SELECT * FROM cto_park")
	if err != nil {
		fmt.Println(err)
	}
	//查看所有列名
	cols, _ := rows.Columns()
	fmt.Println(cols)
	dataRow := make(map[string]interface{})
	for _, col := range cols {
		dataRow[col] = ""
	}
	fmt.Println(dataRow)

	dataRows := make([]interface{}, 10)

	for rows.Next() {
		rows.Scan(dataRow)
		fmt.Println(dataRow)
		fmt.Println(&dataRow)
		dataRows = append(dataRows, dataRow)
	}
	fmt.Println(dataRows)
}
