package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}

func main() {
	a := time.Now().Unix()
	test()
	b := time.Now().Unix()
	c := b - a
	fmt.Println(c)

}
