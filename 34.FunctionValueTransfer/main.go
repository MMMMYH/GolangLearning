/*
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2023-03-25 22:28:27
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2023-03-25 22:48:30
 * @FilePath: \go_code\34 函数\main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
)

func swap(n1 *int, n2 *int) {
	t := *n1
	*n1 = *n2
	*n2 = t
} //值传递对主函数的数值没有影响

func main() {
	c := 10
	d := 20
	swap(&c, &d)
	fmt.Println("反转后为", c, d)
}
