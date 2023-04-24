/*
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2023-03-26 10:08:11
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2023-03-26 10:12:11
 * @FilePath: \go_code\35 init函数\main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"

var age = test()

func test() int {
	fmt.Println("test()")
	return 90
}

func init() {
	fmt.Println("init()")
}

func main() {
	fmt.Println("main()")
}

//init函数可以完成初始化的工作
//先执行定义全局变量，再执行init（），再执行main（）
