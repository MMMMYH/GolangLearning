package main

import "fmt"

func main() {
	//1.for循环  2.for-range
	var a = [...]int{5, 9, 6, 4, 1, 2, 3, 6, 4, 8, 6, 4}
	for index, value := range a {
		fmt.Printf("index=%v value=%v\n", index, value)
	}
	for _, value := range a {
		fmt.Printf("value=%v\n", value)
	}
}
