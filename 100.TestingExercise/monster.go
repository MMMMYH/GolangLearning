package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Monster struct {
	Name   string
	Age    int
	Skills string
}

func (m *Monster) Store() bool {
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println("序列化", err)
		return false
	}
	filePath := "d:/monster.ser"
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("创建文件", err)
		return false
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("写入文件", err)
		return false
	}
	return true
}

func (m *Monster) ReStore() bool {
	filePath := "d:/monster.ser"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("打开文件", err)
		return false
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("读取文件", err)
		return false
	}
	err = json.Unmarshal(data, m)
	if err != nil {
		fmt.Println("反序列化", err)
		return false
	}
	return true
}
