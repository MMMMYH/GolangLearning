/*
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2023-03-25 17:10:35
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2023-03-25 17:13:47
 * @FilePath: \go_code\31 递归2\main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"

func hanshu(n int) int {
	if n == 1 {
		return 3
	} else {
		return 2*hanshu(n-1) + 1
	}
}

func main() {
	var a int
	fmt.Println("请输入n的值")
	fmt.Scanln(&a)
	fmt.Println(hanshu(a))
}
