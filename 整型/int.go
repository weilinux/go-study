package main

import "fmt"

func main () {
	i1 := 10
	// 十进制
	fmt.Printf("%d\n", i1)
	fmt.Printf("%b\n", i1)	// 把十进制数转换成二进制
	fmt.Printf("%o\n", i1)	// 把十进制数转换成八进制
	fmt.Printf("%x\n", i1)	// 把十进制数转换成十六进制

	i2 := 077
	// 八进制
	fmt.Printf("%d\n", i2)

	i3 := 0x1234567
	// 十六进制
	fmt.Printf("%d\n", i3)

	i4 := int8(127)
	// 8位整型 （-128 ~ 127） int32，int64 分别是32位整型，64位整型
	fmt.Printf("%d\n", i4)

	// 查看变量类型
	fmt.Printf("%T\n", i3)
}