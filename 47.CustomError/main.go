package main

import (
	"errors"
	"fmt"
)

func readConf(name string) (err error) {
	if name == "config.ini" {
		return nil
	} else {
		//返回一个自定义错误
		return errors.New("读取文件错误")
	}
}

func test() {
	err := readConf("config.ini")
	if err != nil {
		//如果读取失败发送错误，就输出这个错误，终止程序
		panic(err)
	}
	fmt.Println("test()继续执行")
}

func main() {
	test()
	fmt.Println("main()下面的代码")
}
