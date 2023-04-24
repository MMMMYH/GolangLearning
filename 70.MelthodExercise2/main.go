package main

import "fmt"

type Student struct {
	Name   string
	Gender string
	Age    int
	id     int
	Score  float64
}

type Box struct {
	length float64
	width  float64
	height float64
}

type Visitor struct {
	Name string
	Age  int
}

func (v *Visitor) showPrice() {
	if v.Age > 18 {
		fmt.Println("门票价格:20")
	} else {
		fmt.Println("门票免费")
	}
}
func (b *Box) volume() float64 {
	fmt.Printf("请输入长度:")
	fmt.Scanln(&b.length)
	fmt.Printf("请输入宽度:")
	fmt.Scanln(&b.width)
	fmt.Printf("请输入高度:")
	fmt.Scanln(&b.height)
	return b.length * b.width * b.height
}
func (s *Student) say() string {
	infoStr := fmt.Sprintf("姓名:%s,性别:%s,年龄:%d,学号:%d,成绩:%.2f", s.Name, s.Gender, s.Age, s.id, s.Score)
	return infoStr
}

func main() {
	s1 := Student{
		Name:   "张三",
		Gender: "男",
		Age:    18,
		id:     1,
		Score:  98.5,
	}
	var b1 Box
	var v1 Visitor
	for {
		fmt.Println("请输入名字")
		fmt.Scanln(&v1.Name)
		if v1.Name == "q" {
			break
		}
		fmt.Println("请输入年龄")
		fmt.Scanln(&v1.Age)
		v1.showPrice()
	}
	fmt.Println(s1.say())
	fmt.Println(b1.volume())

}
