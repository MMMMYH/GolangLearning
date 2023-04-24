package main

import "fmt"

func cal(n1 float64, n2 float64, operator byte) float64 {
	var res float64
	switch operator {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
		fmt.Println("输入错误")
	}
	return res
}

func main() {
	a := cal(1.2, 2.3, '+')
	fmt.Println(a)
	var b = 3.1
	var c = 5.6
	var d = '*'
	e := cal(b, c, byte(d))
	fmt.Println(e)
}
