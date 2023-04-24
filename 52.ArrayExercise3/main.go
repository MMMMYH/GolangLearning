package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var a [5]int
	for i := 0; i < len(a); i++ {
		a[i] = rand.Intn(100)
	}
	fmt.Println(a)
	b := 0
	for i := 0; i < len(a)/2; i++ {
		b = a[len(a)-1-i]
		a[len(a)-1-i] = a[i]
		a[i] = b
	}
	fmt.Println(a)
}
