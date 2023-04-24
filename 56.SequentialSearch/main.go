package main

import "fmt"

func main() {
	var a [4]string = [4]string{"白眉鹰王", "金毛狮王", "紫衫龙王", "青翼蝠王"}
	var b string
	fmt.Println("请输入")
	fmt.Scanln(&b)
	// for i := 0; i < len(a); i++ {
	// 	if b == a[i] {
	// 		fmt.Println("包含")
	// 		break
	// 	} else if i == len(a)-1 {
	// 		fmt.Println("不包含")
	// 	}
	// }

	c := -1

	for i := 0; i < len(a); i++ {
		if b == a[i] {
			c = i
			break
		}
	}
	if c != -1 {
		fmt.Println("包含", b)
	} else {
		fmt.Println("不包含", b)
	}
}
