package main

import "fmt"

func main() {
	//输出十句helloworld
	for i := 1; i <= 10; i++ {
		fmt.Println("hello world")
	}
	a := "hello,world!中国"
	b := []rune(a) //将a转换成[]rune切片
	for i := 0; i < len(b); i++ {
		fmt.Printf("%c \n", b[i])
	}
	//for range 遍历
	for index, val := range a {
		fmt.Printf("index=%d,val=%c\n", index, val)
	}
}

//如果字符串有中文，则传统for遍历是错误的 ，因为是按照字节来遍历的
