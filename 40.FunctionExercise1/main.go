package main

import "fmt"

func jinzita(n int) {
	for i := 1; i <= n; i++ {
		for k := 1; k <= n-i; k++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			//if j == 1 || j == 2*i-1 || i == n {
			fmt.Print("*") //不换行
			//} else {
			//fmt.Print(" ")
			//}
		}
		fmt.Println()
	}
}

func main() {
	fmt.Println("请输入")
	var n int
	fmt.Scanln(&n)
	jinzita(n)
}
