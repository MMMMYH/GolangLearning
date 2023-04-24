package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("d:/test.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v", file)
	err = file.Close()
	if err != nil {
		fmt.Println(err)
	}
}
