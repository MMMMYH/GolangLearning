/*
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2023-03-26 10:12:35
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2023-03-26 10:22:49
 * @FilePath: \go_code\36 匿名函数\main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"

var (
	Fun1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func main() {
	res := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println("res=", res)
	//a为函数变量，类型为函数
	a := func(n1 int, n2 int) int {
		return n1 - n2
	}

	res2 := a(10, 20)
	fmt.Println("res2=", res2)

	//全局匿名函数
	res3 := Fun1(4, 6)
	fmt.Println("res3=", res3)
}
