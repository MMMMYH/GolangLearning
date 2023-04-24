package main

import "fmt"

func main() {
	var a [3][5]float64
	var b float64
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			fmt.Printf("请输入第%v个班级的第%v个学生的成绩", i+1, j+1)
			fmt.Scanln(&b)
			a[i][j] = b
		}
	}
	var c, d float64
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			c += a[i][j]
			d += a[i][j]
		}
		fmt.Printf("第%v个班级的平均分为%v", i+1, c/float64(len(a[i])))
		c = 0
	}
	fmt.Printf("所有班级的平均分为%v", d/15)
}
