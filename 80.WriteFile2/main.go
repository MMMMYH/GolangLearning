package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file := "d:/test.txt"
	file1, err := os.OpenFile(file, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file1.Close()
	str := "Hello,北京!\n"
	writer := bufio.NewWriter(file1)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}
