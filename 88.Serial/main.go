package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Birthday string  `json:"birthday"`
	Salary   float64 `json:"salary"`
	Skill    string  `json:"skill"`
}

func testStruct() {
	var monster Monster
	monster.Name = "牛魔王"
	monster.Age = 500
	monster.Birthday = "2000-01-01"
	monster.Salary = 10000
	monster.Skill = "牛魔拳"

	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}

func testMap() {
	a := make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 18
	a["birthday"] = "2000-01-02"
	a["addresss"] = "洪崖洞"

	data, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}

func testSlice() {
	var slice []map[string]interface{}
	m1 := make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = 7
	m1["addresss"] = "北京"
	slice = append(slice, m1)

	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}

func main() {
	testStruct()
	testMap()
	testSlice()
}
