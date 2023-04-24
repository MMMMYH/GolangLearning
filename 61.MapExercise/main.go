package main

import "fmt"

func main() {
	studentMap := make(map[string]map[string]string)
	studentMap["student1"] = make(map[string]string, 4)
	studentMap["student1"]["name"] = "张三"
	studentMap["student1"]["age"] = "18"
	studentMap["student1"]["sex"] = "男"
	studentMap["student1"]["address"] = "北京"
	studentMap["student2"] = make(map[string]string, 3)
	studentMap["student2"]["name"] = "李四"
	studentMap["student2"]["age"] = "20"
	studentMap["student2"]["sex"] = "女"
	studentMap["student2"]["address"] = "上海"
	studentMap["student3"] = make(map[string]string, 2)
	studentMap["student3"]["name"] = "王五"
	studentMap["student3"]["age"] = "22"
	studentMap["student3"]["sex"] = "男"
	studentMap["student3"]["address"] = "广州"
	fmt.Println(studentMap)
	for key, value := range studentMap {
		for key2, value2 := range value {
			fmt.Println(key, key2, value2)
		}
	}
}
