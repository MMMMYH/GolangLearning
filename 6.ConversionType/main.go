package main

import "fmt"

func main() {
	var a = 100
	var b = int(a)
	var c = float32(a)
	fmt.Printf("%T %T %T\n", a, b, c)
	fmt.Println(a, b, c)
	fmt.Printf("a=%v b=%v c=%v", a, b, c)
}
