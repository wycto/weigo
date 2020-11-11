package controller

import (
	"fmt"
)

type IndexController struct {
}

func (receiver *IndexController) Index() {
	fmt.Println("welcome weigo")
}
