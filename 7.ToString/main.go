package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a = 99
	var b = 23.456
	var c = true
	var d byte = 'h'
	fmt.Printf("%T %T %T %T\n", a, b, c, d)
	var e string
	e = fmt.Sprintf("%d", a)
	fmt.Printf("%T %s\n", e, e)
	e = fmt.Sprintf("%f", b)
	fmt.Printf("%T %s\n", e, e)
	e = fmt.Sprintf("%t", c)
	fmt.Printf("%T %s\n", e, e)
	e = fmt.Sprintf("%q", d)
	fmt.Printf("%T %s\n", e, e)

	e = strconv.FormatInt(int64(a), 10)
	fmt.Printf("%T %s\n", e, e)

	e = strconv.FormatFloat(b, 'f', 10, 64) //f是格式 10表示小数点后保留十位 64 表示小数是float64
	fmt.Printf("%T %s\n", e, e)

	e = strconv.FormatBool(c)
	fmt.Printf("%T %s\n", e, e)

}
