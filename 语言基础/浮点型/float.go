package main

import "fmt"

func main() {
	f1 := 1.123456
	fmt.Printf("%T\n", f1) // 默认Go语言中的小数都是float64类型

	f2 := float32(1.123456)
	fmt.Printf("%T\n", f2) // 声明float32类型
}
