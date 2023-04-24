package main

import "fmt"

// func main() {
// 	for i := 1; i <= 100; i++ {
// 		if i%2 == 0 {
// 			continue
// 		}
// 		fmt.Println(i)
// 	}

// }

func main() {
	var a int
	b := 0
	c := 0
	for {
		fmt.Println("请输入数字")
		fmt.Scanln(&a)
		if a == 0 {
			break
		}
		if a > 0 {
			b++
			continue
		}
		c++

	}
	fmt.Println("正数有", b, "个", "负数有", c, "个")

}
