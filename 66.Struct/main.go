package main

import "fmt"

type Cat struct {
	Name  string
	Age   int
	Color string
}

type person struct {
	Name   string
	Age    int
	Scores [5]float64
	ptr    *int
	slice  []int
	map1   map[string]string
}

func main() {
	cat1 := Cat{
		Name:  "小白",
		Age:   1,
		Color: "白色",
	}
	cat2 := Cat{
		Name:  "小黑",
		Age:   2,
		Color: "黑色",
	}
	cat3 := Cat{
		Name:  "小黄",
		Age:   3,
		Color: "黄色",
	}
	fmt.Println(cat1.Name, cat1.Age, cat1.Color)
	fmt.Println(cat2.Name, cat2.Age, cat2.Color)
	fmt.Println(cat3.Name, cat3.Age, cat3.Color)

	var p1 person
	p1.Name = "小明"
	p1.Age = 18
	p1.Scores = [5]float64{1, 2, 3, 4, 5}
	p1.ptr = new(int)
	*p1.ptr = 100
	p1.slice = make([]int, 5)
	p1.slice = []int{1, 2, 3, 4, 5}
	p1.map1 = make(map[string]string)
	p1.map1 = map[string]string{"name": "小明", "age": "18"}
	fmt.Println(p1)
	var p2 *person = new(person)
	(*p2).Name = "小红"
	(*p2).Age = 20
	p2.Scores = [5]float64{1, 2, 3, 4, 5} //简化写法
	var p3 *person = &person{"mark", 20, [5]float64{1, 2, 3, 4, 5}, new(int), []int{1, 2, 3, 4, 5}, make(map[string]string)}
	p3.Name = "小明"
	p3.Age = 18
	fmt.Println(p2, p3)
}
