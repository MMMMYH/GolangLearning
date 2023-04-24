package main

import "fmt"

type A struct {
	name string
	age  int
}

type B struct {
	A
	name string
}

type C struct {
	A
}

func main() {
	var b B
	var c C
	b.name = "b"
	b.age = 10
	b.A.name = "a"
	c.A.name = "c"
	c.A.age = 20
	fmt.Println(b)
	fmt.Println(b.name)
	fmt.Println(b.A.name)
	fmt.Println(c)
	fmt.Println(c.A.name)
	fmt.Println(c.name)
}
