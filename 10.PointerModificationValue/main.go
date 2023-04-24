package main

import "fmt"

func main() {
	var a = 9
	fmt.Printf("%v\n", &a)
	var b *int
	b = &a
	*b = 10
	fmt.Println(a)
}
