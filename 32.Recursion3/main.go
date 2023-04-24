/*
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2023-03-25 17:16:12
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2023-03-25 17:32:49
 * @FilePath: \go_code\32 递归3\main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"

func shengyutaozi(n int) int {
	if n > 10 || n < 1 {
		fmt.Println("输入错误")
		return 0
	}
	if n == 10 {
		return 1
	} else {
		return (shengyutaozi(n+1) + 1) * 2
	}
}

func main() {
	var a int
	fmt.Println("请输入需要查看第几天的剩余桃子数量")
	fmt.Scanln(&a)
	fmt.Println("第", a, "天剩余", shengyutaozi(a), "个桃子")
}
