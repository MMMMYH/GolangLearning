package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a = "true"
	var b bool
	b, _ = strconv.ParseBool(a)
	fmt.Printf("%T %v\n", b, b)
	var c = "123456789"
	var d int64
	d, _ = strconv.ParseInt(c, 10, 64)
	fmt.Printf("%T %v\n", d, d)
	var e = "123.456"
	var f float64
	f, _ = strconv.ParseFloat(e, 64)
	fmt.Printf("%T %v", f, f)
}
