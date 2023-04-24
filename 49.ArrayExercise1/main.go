package main

import "fmt"

func main() {
	var a [5]float64
	for i := 0; i < len(a); i++ {
		fmt.Println("请输入")
		var b float64
		fmt.Scanln(&b)
		a[i] = b
	}
	for c := 0; c < len(a); c++ {
		fmt.Printf("第%v个成绩是%v\n", c+1, a[c])
	}

}
