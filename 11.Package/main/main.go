package main

import (
	"Basic/11.Package/model"
	"fmt"
)

func main() {
	a := model.Cal(1.2, 2.3, '+')
	fmt.Println(a)
	var b = 3.1
	var c = 5.6
	var d = '*'
	e := model.Cal(b, c, byte(d))
	fmt.Println(e)
}
