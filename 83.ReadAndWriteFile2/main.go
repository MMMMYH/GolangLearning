package main

import (
	"fmt"
	"os"
)

func main() {
	file1 := "d:/test.txt"
	file2 := "d:/test2.txt"
	concent, err := os.ReadFile(file1)
	if err != nil {
		fmt.Println(err)
		return
	}
	line, err1 := os.OpenFile(file2, os.O_WRONLY|os.O_CREATE, 0666)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	defer line.Close()
	str := string(concent)
	line.WriteString(str)
}
