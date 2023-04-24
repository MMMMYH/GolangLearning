package main

import "fmt"

func main() {
	var a int
	fmt.Scanln(&a)
	if a == 100 {
		fmt.Println("奖励一辆BMW")
	} else if a > 80 && a <= 99 {
		fmt.Println("奖励一台iphone7plus")
	} else if a >= 60 && a <= 80 {
		fmt.Println("奖励一个ipad")
	} else {
		fmt.Println("什么都没有")
	}
}
