package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file := "d:/test.txt"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(content))
}
