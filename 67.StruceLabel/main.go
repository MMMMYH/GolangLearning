package main

import (
	"encoding/json"
	"fmt"
)

type Students struct {
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Score float64 `json:"score"`
}

func main() {
	var s1 Students = Students{"小明", 18, 3.14}
	jsonStudents, _ := json.Marshal(s1)
	fmt.Println(string(jsonStudents))
}
