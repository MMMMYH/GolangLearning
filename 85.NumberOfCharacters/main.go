package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type CharCount struct {
	ChCount    int //英文
	NumCount   int //数字
	SpaceCount int //空格
	OtherCount int //其他

}

func main() {
	fileName := "d:/test3.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()
	var count CharCount
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		//rune := []rune(str)
		for _, v := range str {
			switch {
			case v >= 'a' && v <= 'z' || v >= 'A' && v <= 'Z':
				count.ChCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			case v == ' ':
				count.SpaceCount++
			default:
				count.OtherCount++
			}
		}
	}
	fmt.Println(count)
}
