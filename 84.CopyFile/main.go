package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CopyFile(dstFileName, srcFileName string) (written int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println(err)
	}
	reader := bufio.NewReader(srcFile)
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	writer := bufio.NewWriter(dstFile)
	defer srcFile.Close()
	defer dstFile.Close()
	return io.Copy(writer, reader)
}
func main() {
	srcFile := "d:/myh.jpg"
	dstFile := "d:/myh2.jpg"
	_, err := CopyFile(dstFile, srcFile)
	if err != nil {
		fmt.Println(err)
		return
	}
}
