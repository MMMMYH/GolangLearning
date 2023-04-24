package main

import (
	"fmt"
)

func main() {
	cities := make(map[int]string)
	cities[1] = "北京"
	cities[2] = "上海"
	cities[3] = "广州"
	cities[4] = "深圳"
	fmt.Println(cities)
	value, ok1 := cities[1]
	if ok1 {
		fmt.Printf("有1 key的值为%v\n", value)
	} else {
		fmt.Println("没有1")
	}
	cities[1] = "杭州"
	cities[5] = "成都"
	fmt.Println(cities)
	delete(cities, 1)
	value, ok2 := cities[1]
	if ok2 {
		fmt.Printf("有1 key的值为%v\n", value)
	} else {
		fmt.Println("没有1")
	}
	delete(cities, 6)
	fmt.Println(cities)
	cities = make(map[int]string)
	fmt.Println(cities)
}
