package main

import "fmt"

func main() {
	s1 := "指针"
	a1 := [...]int{1, 2, 3, 4, 5}

	// 获取指针地址
	p1 := &s1
	p2 := &a1
	fmt.Printf("%T => %p\n", p1, p1)
	fmt.Printf("%T => %p\n", p2, p2)

	// 通过指针取值
	v1 := *p1
	fmt.Println(v1)

	// 声明指针类型的变量 nil
	var p3 *int
	// new()创建匿名变量申请内存空间并返回生成的指针地址
	var p4 = new(int)
	*p4 = 100
	p3 = p4
	fmt.Println(p3)
	fmt.Println(*p4)
}
