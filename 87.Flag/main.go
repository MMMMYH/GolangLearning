package main

import (
	"flag"
	"fmt"
)

func main() {
	var user string
	var pwd string
	var host string
	var port int

	flag.StringVar(&user, "u", "", "用户名")
	flag.StringVar(&pwd, "pwd", "", "密码")
	flag.StringVar(&host, "h", "loaclhost", "主机名")
	flag.IntVar(&port, "port", 3306, "端口号")
	flag.Parse() //解析参数
	fmt.Println(user, pwd, host, port)
}
