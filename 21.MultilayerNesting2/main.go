package main

import "fmt"

func main() {
	var a int //行数
	fmt.Println("请输入行数")
	fmt.Scanln(&a)
	for i := 1; i <= a; i++ {
		for k := 1; k <= a-i; k++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			if j == 1 || j == 2*i-1 || i == a {
				fmt.Print("*") //不换行
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
