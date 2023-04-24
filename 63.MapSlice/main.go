package main

import "fmt"

func main() {
	var students []map[string]string
	students = make([]map[string]string, 3)
	if students[0] == nil {
		students[0] = make(map[string]string)
		students[0]["name"] = "张三"
		students[0]["age"] = "18"
		students[0]["sex"] = "男"
		students[0]["address"] = "北京"
	}
	if students[1] == nil {
		students[1] = make(map[string]string)
		students[1]["name"] = "李四"
		students[1]["age"] = "20"
		students[1]["sex"] = "女"
		students[1]["address"] = "上海"
	}
	if students[2] == nil {
		students[2] = make(map[string]string)
		students[2]["name"] = "王五"
		students[2]["age"] = "22"
		students[2]["sex"] = "男"
		students[2]["address"] = "广州"
	}
	newStudent := make(map[string]string)
	newStudent["name"] = "赵六"
	newStudent["age"] = "25"
	newStudent["sex"] = "女"
	newStudent["address"] = "深圳"
	students = append(students, newStudent)
	fmt.Println(students)
}
