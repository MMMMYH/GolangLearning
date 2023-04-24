package main

import (
	"fmt"
	"time"
)

func main() {
	//获取当前时间
	a := time.Now()
	fmt.Printf("now=%v\ntype=%T\n", a, a)

	//获取其他日期信息
	fmt.Println(a.Year())
	fmt.Println(time.Now().Year())
	fmt.Println(a.Month())
	fmt.Println(time.Now().Month())
	fmt.Println(int(a.Month()))
	fmt.Println(int(time.Now().Month()))
	fmt.Println(a.Day())
	fmt.Println(time.Now().Day())
	fmt.Println(a.Hour())
	fmt.Println(time.Now().Hour())
	fmt.Println(a.Minute())
	fmt.Println(time.Now().Minute())
	fmt.Println(a.Second())
	fmt.Println(time.Now().Second())

	//格式化 日期时间
	fmt.Printf("当前年月日%02d-%02d-%02d %02d:%02d:%02d\n", a.Year(), a.Month(), a.Day(), a.Hour(), a.Minute(), a.Second())
	b := fmt.Sprintf("当前年月日%02d-%02d-%02d %02d:%02d:%02d\n", a.Year(), a.Month(), a.Day(), a.Hour(), a.Minute(), a.Second())
	fmt.Println(b)
	fmt.Printf(a.Format("2006/01/0215:04:05"))
	fmt.Println()
	fmt.Printf(a.Format("2006/01/02"))
	fmt.Println()
	fmt.Printf(a.Format("15:04:05"))
	fmt.Println()

	//时间常量 在程序中用于获取指定时间单位的时间，比如想得到100毫秒 100*time.Millisecond
	// const (
	// 	Nanosecond  Duration = 1                  //纳秒
	// 	Microsecond          = 1000 * Nanosecond  //微秒
	// 	Millisecond          = 1000 * Microsecond //毫秒
	// 	Second               = 1000 * Millisecond //秒
	// 	Minute               = 60 * Second        //分钟
	// 	Hour                 = 60 * Minute
	// )

	//休眠
	c := 0
	for {
		c++
		fmt.Println(c)
		time.Sleep(time.Second / 1000)
		if c == 10 {
			break
		}
	}

	//获取当前unix时间戳和unixnano时间戳 作用是获取随机数字
	//从时间点1970/1/1 UTC到时间点t所经过的时间
	fmt.Printf("unix时间戳=%v\nunixnano时间戳=%v", a.Unix(), a.UnixNano())

}
