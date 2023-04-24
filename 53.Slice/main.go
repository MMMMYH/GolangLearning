package main

import "fmt"

func main() {
	var intArr [5]int = [...]int{11, 22, 33, 44, 55}
	slice := intArr[1:3] //={22,33}
	fmt.Println("intArr=", intArr)
	fmt.Println("slice的元素是=", slice)
	fmt.Println("slice的元素个数是=", len(slice))
	fmt.Println("slice的容量是", cap(slice))

	var slice1 []int = make([]int, 3, 6)
	fmt.Println(slice1)

	var slice2 []int = []int{1, 2, 3, 4, 5}
	fmt.Println("slice2的元素是=", slice2)
	fmt.Println("slice2的元素个数是=", len(slice2))
	fmt.Println("slice2的容量是", cap(slice2))

	for i := 0; i < len(slice2); i++ {
		fmt.Println(slice2[i])
	}
	for i, v := range slice2 {
		fmt.Printf("i=%v v=%v\n", i, v)
	}

	var slice3 []int = []int{1, 2, 3}
	slice3 = append(slice3, 4, 5, 6)
	slice3 = append(slice3, slice3...)
	fmt.Println(slice3)

	var slice4 []int = []int{1, 2, 3, 4, 5}
	var slice5 = make([]int, 10)
	fmt.Println(slice5)
	copy(slice5, slice4)
	fmt.Println(slice5)

}
