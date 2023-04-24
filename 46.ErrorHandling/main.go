package main

import (
	"fmt"
)

func test() {
	//defer+recover
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err=", err)
		}
	}()
	a := 10
	b := 0
	fmt.Println(a / b)
}

func main() {
	test()
	fmt.Println("下面代码")
}

//defer panic recover来处理
