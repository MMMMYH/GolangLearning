package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("d:/test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n') //读取一行
		if err == io.EOF {
			break
		}
		fmt.Print(line)
	}
	fmt.Println("文件读取完毕")
}
