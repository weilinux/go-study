package main

import (
	"fmt"
	"unicode"
)

func main() {
	/*
	* 一种是 uint8 类型，或者叫 byte 型，代表了 ASCII 码的一个字符，byte 类型是 uint8 的别名
	* 另一种是 rune 类型，代表一个 UTF-8 字符，当需要处理中文、日文或者其他复合字符时，则需要用到 rune 类型。rune 类型等价于 int32 类型。
	 */

	//在 ASCII 码表中，A 的值是 65，使用 16 进制表示则为 41
	var ch1 byte = 65
	var ch2 byte = '\x41'

	fmt.Println(ch1)
	fmt.Println(ch2)

	ch3 := "hello长沙市0731（邮编）"
	for i := 0; i < len(ch3); i++ {
		fmt.Printf("%v--%c\n", ch3[i], ch3[i])
	}
	fmt.Println()

	for index, value := range ch3 {
		fmt.Println(index)
		fmt.Printf("%v--%c\n", value, value)
	}

	// 类型转换
	ch4 := 10
	var ch5 float64
	ch5 = float64(ch4)
	fmt.Printf("%v---%T\n", ch4, ch4)
	fmt.Printf("%v---%T\n", ch5, ch5)

	// 获取"hello长沙"中汉字的数量
	var cheseNumber int
	for _, value := range ch3 {
		if unicode.Is(unicode.Han, value) {
			cheseNumber++
		}
	}
	fmt.Print(cheseNumber)
}
