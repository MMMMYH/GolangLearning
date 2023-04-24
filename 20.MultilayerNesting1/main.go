package main

import (
	"fmt"
)

func main() {
	var a float64 = 0.00
	var b float64 = 0.00
	var d float64
	var e int = 0
	fmt.Println("请输入班级数量")
	var m int
	fmt.Scanln(&m)
	fmt.Println("请输入班级中学生的数量")
	var n int
	fmt.Scanln(&n)
	var c int = 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			fmt.Println("请输入第", i, "个班的第", j, "位学生的成绩")
			fmt.Scanln(&d)
			if d >= 60 {
				c += 1
			}
			a += d
		}

		fmt.Println("第", i, "个班的平均分为", a/float64(n), "第", i, "个班的及格人数为", c)
		b += a
		e += c
		a = 0
		c = 0
	}
	fmt.Println("所有班级总成绩为", b, "所有班级平均分", b/float64(m*n), "所有班级及格人数为", e)
}
