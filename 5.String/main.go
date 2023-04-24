package main

import "fmt"

func main() {
	var a = "北京长城"
	fmt.Println(a)
	fmt.Printf("%T", a) //类型
	//用反引号``
	var b = `package main

	import (
		"fmt"
		"unsafe"
	)
	
	func main() {
		var n1 = "tom"
		var a1 byte = 'a'
		var a2 byte = '0'
		var a3 = "a"
		var a4 = "0"
		var a5 = '北'
		fmt.Printf("n1的类型是 %T \n", n1) //查看类型"%T"
		fmt.Println("n1=", n1)
		fmt.Printf("n1所占的字节数是%d\n", unsafe.Sizeof(n1)) //查看所占字节数大小"%d"
		fmt.Printf("a1的类型是 %T \n", a1)                 //查看类型"%T"
		fmt.Println("a1=", a1)
		fmt.Printf("a1=%c\n", a1)      //按照字符形式输出"%c"
		fmt.Printf("a2的类型是 %T \n", a2) //查看类型"%T"
		fmt.Println("a2=", a2)
		fmt.Printf("a3的类型是 %T \n", a3) //查看类型"%T"
		fmt.Println("a3=", a3)
		fmt.Printf("a4的类型是 %T \n", a4) //查看类型"%T"
		fmt.Println("a4=", a4)
		fmt.Printf("a5的类型是%T\n", a5)
		fmt.Println("a5=", a5)
		fmt.Printf("a5=%c\n", a5)
	
	}`
	fmt.Println(b)
	c := "hello" + "world" +
		"hello" + "world" +
		"hello" + "world" +
		"hello" + "world" +
		"hello" + "world"
	fmt.Println(c)
}
