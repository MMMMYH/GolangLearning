/*
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2023-03-25 19:10:07
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2023-03-25 19:13:57
 * @FilePath: \go_code\33 调用函数内变量\main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"

func test(n1 *int) {
	*n1 = *n1 + 10
	fmt.Println("test() n1=", *n1)
}
func main() {
	num := 20
	test(&num)
	fmt.Println("main() num=", num)
}
