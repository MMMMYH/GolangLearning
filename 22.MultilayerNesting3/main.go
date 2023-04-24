package main

import "fmt"

func main() {
	var a int = 20
	for i := 1; i <= a; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j, "*", i, "=", i*j, " ") //fmt.Printf("%v * %v = %v \t",j,i,j*i)
		}
		fmt.Println()
	}
}
