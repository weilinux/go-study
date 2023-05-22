package main

import "fmt"

func main() {
	n := 100

	// 类型
	fmt.Printf("%T\n", n)
	// 值
	fmt.Printf("%v\n", n)
	// 二进制
	fmt.Printf("%b\n", n)
	// 八进制
	fmt.Printf("%o\n", n)
	// 十进制
	fmt.Printf("%d\n", n)
	// 十六进制
	fmt.Printf("%x\n", n)

	s := "hello"
	// 字符串
	fmt.Printf("%s\n", s)
	// 添加变量描述符 string类型会加上双引号（“”）
	fmt.Printf("%#v\n", n)
	fmt.Printf("%#v\n", s)
}
