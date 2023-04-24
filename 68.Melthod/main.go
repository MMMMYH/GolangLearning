package main

import "fmt"

type A struct {
	Num int
}

func (a A) test() {
	fmt.Println(a.Num)
}

func main() {
	b := A{5}
	b.test()
}
