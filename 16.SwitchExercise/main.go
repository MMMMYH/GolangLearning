package main

import (
	"fmt"
)

//1
func main() {
	var a string
 	fmt.Println("请输入字母a/b/c/d/e")
 	fmt.Scanln(&a)
 	switch a {
 	case "a":
 		fmt.Println("A")
 	case "b":
 		fmt.Println("B")
 	case "c":
 		fmt.Println("C")
 	case "d":
 		fmt.Println("D")
 	case "e":
 		fmt.Println("E")
 	default:
		fmt.Println("输入错误")
 	}
 }

// 2
// func main() {
// 	var a int
// 	fmt.Println("请输入成绩")
// 	fmt.Scanln(&a)
// 	switch {
// 	case a >= 0 && a < 60:
// 		fmt.Println("不及格")
// 	case a >= 60 && a <= 100:
// 		fmt.Println("及格")
// 	default:
// 		fmt.Println("输入错误")
// 	}
// }

// 3
// func main() {
// 	var a int
// 	fmt.Println("请输入月份")
// 	fmt.Scanln(&a)
// 	switch {
// 	case a == 3 || a == 4 || a == 5:
// 		fmt.Println("春季")
// 	case a == 6 || a == 7 || a == 8:
// 		fmt.Println("夏季")
// 	case a == 9 || a == 10 || a == 11:
// 		fmt.Println("秋季")
// 	case a == 12 || a == 1 || a == 2:
// 		fmt.Println("冬季")
// 	default:
// 		fmt.Println("输入错误")
// 	}
// }

//4
