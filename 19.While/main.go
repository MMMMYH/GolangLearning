package main

import "fmt"

//go语言没有 ，只能用for

//for实现while
// func main() {

// 	i := 1
// 	for {
// 		if i > 10 {
// 			break
// 		}
// 		fmt.Println("hello,world", i)
// 		i++
// 	}
// }

//for实现do...while
func main() {
	i := 1
	for {
		fmt.Println("hello,world", i)
		i++
		if i > 10 {
			break
		}
	}
}
