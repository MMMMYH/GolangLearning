package main

import (
	"Basic/71.Package/model"
	"fmt"
)

func main() {
	p := model.Newperson("smith")
	p.SetAge(20)
	p.SetSal(10000)
	fmt.Println(p)
}
