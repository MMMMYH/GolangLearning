package main

import "fmt"

func main() {
	var a map[string]string
	a = make(map[string]string, 10)
	a["name1"] = "张三"
	a["name2"] = "李四"
	a["name3"] = "王五"
	a["name4"] = "赵六"
	a["name5"] = "孙七"
	fmt.Println(a)
	for k, v := range a {
		fmt.Println(k, v)
	}
}
