package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var a int = 123456
	var b string = "Hello,World!北京"
	var bb string = "123456789"
	var bbb byte = 97

	//统计字符串长度（UTF-8  字母数字占1个字节 汉字3个字节
	c := len(b)
	fmt.Println(c)

	//字符串遍历，处理有中文的问题
	for i := 0; i < len(b); i++ {
		fmt.Printf("字符=%c\n", b[i])
	} //会出现乱码
	d := []rune(b)

	for i := 0; i < len(d); i++ {
		fmt.Printf("字符=%c\n", d[i])
	}
	fmt.Println(d)

	//字符串转整数 整数转字符串
	e, _ := strconv.Atoi(bb)
	f := strconv.Itoa(a)
	fmt.Println(e, f)
	fmt.Printf("e的类型是%T\n", e)
	fmt.Printf("f的类型是%T\n", f)

	//字符串转byte byte转字符串
	g := []byte(b)
	fmt.Printf("g的类型是%T\n", g)
	fmt.Println(g)
	h := string([]byte{bbb})
	fmt.Printf("h的类型是%T\n", h)
	fmt.Printf("%v\n", h)

	//10进制转2/8/16进制 返回对应字符串
	i := strconv.FormatInt(int64(a), 2)
	j := strconv.FormatInt(int64(a), 8)
	k := strconv.FormatInt(int64(a), 16)
	fmt.Println(i, j, k)

	//c查找子串是否在指定字符串中 ，返回bool
	fmt.Println(strings.Contains(b, "北京"))
	fmt.Println(strings.Contains(b, "上海"))

	//统计一个字符串有几个指定的子串
	fmt.Println(strings.Count(b, "o"))
	fmt.Println(strings.Count(b, "上海"))

	//不区分大小写的字符串比较    区分大小写字符串比较==
	l := strings.EqualFold(b, "HELLO,WORLD!北京")
	fmt.Println(l)
	fmt.Println(b == "HELLO,WORLD!北京")

	//返回子串在字符串第一个出现的index值，如果没有返回-1
	fmt.Println(strings.Index(b, "ll"))
	fmt.Println(strings.Index(b, "lla"))

	//返回子串在字符串最后一次出现的index值，如果没有返回-1
	fmt.Println(strings.LastIndex(b, "l"))
	fmt.Println(strings.LastIndex(b, "oo"))

	//将指定子串替换成另一个子串 n可以指定希望换一个 n=-1表示换全部
	fmt.Println(strings.Replace("Hello,World!北京，北京。", "北京", "上海", -1))
	fmt.Println(strings.Replace("Hello,World!北京，北京。", "北京", "上海", 1))
	fmt.Println(strings.Replace("Hello,World!北京，北京。", "北京", "上海", 2))
	fmt.Println(strings.Replace("Hello,World!北京，北京。", "北京", "上海", 3))

	//按照指定的某个字符为分割标志，将一个字符串拆分为字符串数组
	fmt.Println(strings.Split("aabacadaea", "a"))

	//字母大小写转换
	fmt.Println(strings.ToLower(b))
	fmt.Println(strings.ToUpper(b))

	//去除字符串左右两端空格
	fmt.Println(strings.TrimSpace("     hello     "))
	fmt.Println(strings.TrimSpace(b))

	//去除字符串左右两端指定的字符
	fmt.Println(strings.Trim("!hello!", "!"))
	fmt.Println(strings.Trim("!!hello!!", "!"))
	fmt.Println(strings.Trim("hello!", "!"))

	//去除左边/右边的指定字符
	fmt.Println(strings.TrimLeft("!hello!", "!"))
	fmt.Println(strings.TrimRight("!hello!", "!"))

	//判断字符串是否以指定字符串开头
	fmt.Println(strings.HasPrefix(b, "h"))
	fmt.Println(strings.HasPrefix(b, "H"))
	fmt.Println(strings.HasPrefix(b, "hello"))
	fmt.Println(strings.HasPrefix(b, "北京"))

	//判断字符串是否以指定字符串结尾
	fmt.Println(strings.HasSuffix(b, "北京"))
}
