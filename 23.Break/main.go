package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a := 1
	for {
		//rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100) + 1
		if n == 99 {
			fmt.Println(n)
			fmt.Println("随机次数为", a)
			break
		} else {
			a += 1
			fmt.Println(n)
		}
	}

}
