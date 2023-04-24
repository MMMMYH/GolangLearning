package main

import "fmt"

func main() {
	var a int = 2
	fmt.Println(1)
	fmt.Println(2)
	if a < 100 {
		goto label1
	}
	fmt.Println(3)
	fmt.Println(4)
label1:
	fmt.Println(5)
	fmt.Println(6)
	fmt.Println(7)
	fmt.Println(8)
	fmt.Println(9)
	fmt.Println(10)
}
