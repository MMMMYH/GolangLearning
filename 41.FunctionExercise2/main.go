package main

import "fmt"

func chengfabiao(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j, "*", i, "=", i*j, " ") //fmt.Printf("%v * %v = %v \t",j,i,j*i)
		}
		fmt.Println()
	}
}

func main() {
	fmt.Println("请输入")
	var a int
	fmt.Scanln(&a)
	chengfabiao(a)
}
