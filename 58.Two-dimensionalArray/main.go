package main

import "fmt"

func main() {
	var a [4][6]int
	a[1][2] = 1
	a[2][1] = 2
	a[2][3] = 3
	var b [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			fmt.Print(a[i][j], " ")
		}
		fmt.Println()
	}
	for i, v := range b {
		for j, v2 := range v {
			fmt.Printf("b[%v][%v]=%v\n", i, j, v2)
		}
	}
}
