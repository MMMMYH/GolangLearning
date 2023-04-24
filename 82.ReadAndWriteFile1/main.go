package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file := "d:/test.txt"
	file1, err := os.OpenFile(file, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file1.Close()
	reader := bufio.NewReader(file1)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(line)
	}

	str := "Hello,北京上海!\n"
	writer := bufio.NewWriter(file1)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}
