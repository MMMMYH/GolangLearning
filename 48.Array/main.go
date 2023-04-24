package main

import "fmt"

// func main() {
// 	a := 3.0
// 	b := 5.0
// 	c := 1.0
// 	d := 3.4
// 	e := 2.0
// 	f := 50.0
// 	g := a + b + c + d + e + f
// 	h := fmt.Sprintf("%.2f", g/6)
// 	fmt.Println(g, h)
// }

//数组是值类型
func main() {
	//定义
	var a [6]float64
	a[0] = 3.0
	a[1] = 5.0
	a[2] = 1.0
	a[3] = 3.4
	a[4] = 2.0
	a[5] = 50.0
	//遍历
	b := 0.0
	for i := 0; i < len(a); i++ {
		b += a[i]
	}
	//求出平均值
	fmt.Printf("%.2f", b/float64(len(a)))

	//定义数组的方式
	var n1 [3]int = [3]int{1, 2, 3}
	fmt.Println("第一种", n1)
	var n2 = [3]int{1, 2, 3}
	fmt.Println("第二种", n2)
	var n3 = [...]int{1, 2, 3}
	fmt.Println("第三种", n3)
	var n4 = [3]int{0: 3, 2: 1, 1: 2}
	fmt.Println("第四种", n4)
	var n5 = [...]string{0: "asd", 2: "fgh", 1: "jkl", 3: ";'"}
	fmt.Println("第五种", n5)
}
