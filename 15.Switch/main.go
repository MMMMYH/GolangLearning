package main

import "fmt"

func main() {
	var a string //byte
	fmt.Println("请输入字符")
	fmt.Scanln(&a) //若为byte 用fmt.Scanf("%a",&a)
	switch a {
	case "a": //若为byte 用''
		fmt.Println("星期一")
	case "b":
		fmt.Println("星期二")
	case "c":
		fmt.Println("星期三")
	case "d":
		fmt.Println("星期四")
	case "e":
		fmt.Println("星期五")
	case "f":
		fmt.Println("星期六")
	case "g":
		fmt.Println("星期日")
	default:
		fmt.Println("输入错误")
	}
}
