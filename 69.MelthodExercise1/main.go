package main

import "fmt"

type MethodUtils struct {
}
type Calculator struct {
	num1 int
	num2 int
}

func (m *MethodUtils) Print() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 8; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (m *MethodUtils) Print2(a int, b int) {
	for i := 0; i < a; i++ {
		for j := 0; j < b; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (m *MethodUtils) Judge(a int) {
	if a%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}
}

func (m *MethodUtils) Print3(a int, b int, c string) {
	for i := 0; i < a; i++ {
		for j := 0; j < b; j++ {
			fmt.Print(c)
		}
		fmt.Println()
	}
}

func (m *MethodUtils) Area(a int, b int) int {
	return a * b
}

func (c *Calculator) Add(a int, b int) int {
	return a + b
}

func (c *Calculator) Sub(a int, b int) int {
	return a - b
}

func (c *Calculator) Mul(a int, b int) int {
	return a * b
}

func (c *Calculator) Div(a int, b int) int {
	return a / b
}

type Result struct {
	Value int
}

func (c *Calculator) Cal(a int, b int, d string) Result {
	if d == "+" {
		return Result{c.Add(a, b)}
	} else if d == "-" {
		return Result{c.Sub(a, b)}
	} else if d == "*" {
		return Result{c.Mul(a, b)}
	} else if d == "/" {
		return Result{c.Div(a, b)}
	}
	return Result{}
}

func main() {
	var m MethodUtils
	var n Calculator
	m.Print()
	fmt.Println("---------------")
	var m1 MethodUtils
	m1.Print2(3, 4)
	fmt.Println("---------------")
	fmt.Println(m.Area(3, 4))
	fmt.Println("---------------")
	m.Judge(0)
	fmt.Println("---------------")
	m.Print3(3, 4, "&")
	fmt.Println("---------------")
	fmt.Println(n.Add(3, 4))
	fmt.Println(n.Sub(3, 4))
	fmt.Println(n.Mul(3, 4))
	fmt.Println(n.Div(3, 4))
	fmt.Println("---------------")
	fmt.Println(n.Cal(3, 4, "+"))
	fmt.Println(n.Cal(3, 4, "-"))
	fmt.Println(n.Cal(3, 4, "*"))
	fmt.Println(n.Cal(3, 4, "/"))

}
