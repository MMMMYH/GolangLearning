package main

import "fmt"

//变量
var (
	a = 1
	b = 2
	c = 3
)

func main() {
	var m int = 5
	var i = 10
	name1, name2 := "tom", "jack"
	var d = 1
	var e = 2
	var f = d + e
	var g = "hello"
	var h = "world"
	var n = g + h
	fmt.Println("m=", m, "i=", i, "name1=", name1, "name2=", name2)
	fmt.Println("a=", a, "b=", b, "c=", c)
	fmt.Println("f=", f)
	fmt.Println("n=", n)
}
