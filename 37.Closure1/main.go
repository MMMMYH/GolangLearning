/*
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2023-03-26 10:26:30
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2023-03-26 10:52:33
 * @FilePath: \go_code\37 闭包\main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"

func AddUpper() func(int) int { //返回函数fun(int)int
	var n int = 10           //闭包开始
	return func(x int) int { //返回一个匿名函数，但匿名函数引用到函数外的n，因此匿名函数和n构成一个整体，构成闭包
		n = n + x
		return n //闭包结束
	}
}

//闭包是类 函数是操作 n是字段 函数和n构成闭包
func main() {
	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))
}

//返回的函数和函数调用的函数外的变量构成闭包
