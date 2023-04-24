package main

import "fmt"

func main() {
	//1
	var a [26]byte
	for i := 0; i < 26; i++ {
		a[i] = 'A' + byte(i)
	}
	for i, v := range a {
		fmt.Printf("index=%v value=%c\n", i, v)
	}

	//2
	b := [...]float64{3.6, 43.8, 9.3, 168.3, 999, 666.3, 9.8}
	var c float64
	d := 0
	for i := 0; i < len(b); i++ {
		if b[i] >= c {
			c = b[i]
			d = i
		}
	}
	fmt.Println(d+1, c)

	//3
	var e float64
	for _, val := range b {
		e += val
	}
	fmt.Println(e, float64(e)/float64(len(b)))
}
