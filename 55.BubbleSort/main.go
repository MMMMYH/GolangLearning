package main

import "fmt"

func BubbleSort(arr *[5]int64) {
	for j := 1; j <= len(arr)-1; j++ {
		for i := 0; i < len(arr)-j; i++ {
			if arr[i] > arr[i+1] {
				a := arr[i+1]
				arr[i+1] = arr[i]
				arr[i] = a
			}
		}
	}
}

func main() {
	var arr [5]int64 = [5]int64{24, 69, 80, 57, 13}
	BubbleSort(&arr)
	fmt.Println(arr)
}
