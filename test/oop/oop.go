package main

import (
	"fmt"
	"wycto/app/index/model"
)

func main() {
	user := model.User{}
	user.Name = "唯一"
	fmt.Println(user.Name)
}
