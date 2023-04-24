package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file := "d:/abc.txt"
	file1, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	str := "hello world\n"
	writer := bufio.NewWriter(file1)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	defer file1.Close()
	writer.Flush() //
}
