package main

import "fmt"

func BinaryFind(arr *[6]int, left int, right int, find int) {
	var mid = (left + right) / 2
	if left == right || left > right {
		fmt.Println("没有找到")
	} else {
		if (*arr)[mid] > find {
			BinaryFind(arr, left, mid-1, find)
		} else if (*arr)[mid] < find {
			BinaryFind(arr, mid+1, right, find)
		} else if (*arr)[mid] == find {
			fmt.Println("找到", find)
		}
	}
}

func main() {
	var a [6]int = [6]int{1, 8, 89, 100, 256, 512}
	var b int
	fmt.Println("请输入")
	fmt.Scanln(&b)
	BinaryFind(&a, 0, len(a), b)

}
