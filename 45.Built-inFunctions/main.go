package main

import "fmt"

//len 求长度 string/array/slice/map/channel
//new 用来分配内存，主要用来分配值类型 int/float32/struct。。。 返回的是指针
//make 用来分配内存 主要用来分配引用类型 channal/map/slice
func main() {
	//new
	a := 100
	fmt.Printf("type=%T\na=%v\n地址=%v\n", a, a, &a)
	b := new(int)
	*b = 100
	fmt.Printf("type=%T\na=%v\n地址=%v\n指针指向的值=%v\n", b, b, &b, *b)
	//改变b的值

}
