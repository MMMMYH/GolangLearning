package main

import "fmt"

// func main() {
// 	var a int
// 	var j int
// 	for i := 1; i <= 100; i++ {
// 		if i%9 == 0 {
// 			a += i
// 			j += 1
// 		}
// 	}
// 	fmt.Println("个数=,总和=", j, a)
// }

func main() {
	var j int = 6
	for i := 0; i <= 6; i++ {
		fmt.Println(i, "+", j, "=", i+j)
		j--
	}
}
