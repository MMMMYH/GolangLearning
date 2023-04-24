package main

import "fmt"

func main() {
	var a string
	var b int
	for i := 1; i <= 3; i++ {
		fmt.Println("请输入用户名")
		fmt.Scanln(&a)
		fmt.Println("请输入密码")
		fmt.Scanln(&b)
		if a == "张无忌" && b == 888 {
			fmt.Println("登陆成功")
			break
		}
		if i == 3 {
			fmt.Println("账户已锁定")
			break
		}
		fmt.Println("还剩", 3-i, "次机会")
	}
}
