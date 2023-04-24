package main

import "fmt"

func main() {
	var i = 10
	//i的地址
	fmt.Println(&i)
	var a *int = &i //a为指针变量 类型是*int 本身的值是&i
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", &a)
	fmt.Printf("%v\n", *a)
}
