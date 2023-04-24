package main

import "fmt"

func main() {
	a := 0
	b := 0
	for i := 1; i <= 100; i++ {
		a += i
		if a >= 20 {
			b = i
			break
		}
	}
	fmt.Println(a, b)
}
