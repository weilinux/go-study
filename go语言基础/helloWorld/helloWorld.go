package main

import "fmt"

const (
	nameLen = 10
	nameCap = 20
)

func main() {
	/*a, b := 13, 26
	fmt.Printf("a %b\n", a)
	fmt.Printf("b %b\n", b)

	a ^= b
	fmt.Printf("a %b\n", a)
	b ^= a
	fmt.Printf("b %b\n", b)
	a ^= b
	fmt.Printf("a %b\n", a)
	fmt.Printf("b %b\n", b)*/

	/*var name string
	var age int

	fmt.Println("请输入姓名")
	_, err := fmt.Scanln(&name)
	if err != nil {
		return
	}

	fmt.Println("请输入年龄")
	_, err = fmt.Scanln(&age)
	if err != nil {
		return
	}

	fmt.Printf("姓名：%v,\n年龄：%v", name, age)*/

	/*
		1 源码[0000 0001] 反码[0000 0001] 补码[0000 0001]

		-1 源码[1000 0001] 反码[1111 1110] 补码[1111 1111]

		-3 源码[1000 0011] 反码[1111 1100] 补码[1111 1101]

		51 源码[0011 0011] 反码[0011 0011] 补码[0011 0011]
		计算前 反码[0000 0011]
		计算后 反码[0000 0011]
		源码[0000 0011]

		-51 源码[1011 0011] 反码[1100 1100] 补码[1100 1101]
		计算前 反码[1111 1100]
		计算后 反码[1111 1011]
		源码[1000 0100]

		i := 51 >> 4
		fmt.Printf("i %b", i)
	*/

	/*v := 152
	fmt.Printf("v: %b\n", v)

	b := v >> 3
	fmt.Printf("v: %b\n", v)
	fmt.Printf("b: %b\n", b)
	fmt.Printf(strconv.Itoa(b & 1))*/

	/*var s []int // len(s) == 0, s == nil
	fmt.Println(len(s), cap(s))

	s = nil // len(s) == 0, s == nil
	fmt.Println(len(s), cap(s))

	s = []int(nil) // len(s) == 0, s == nil
	fmt.Println(len(s), cap(s))

	s = []int{} // len(s) == 0, s != nil
	fmt.Println(len(s), cap(s))

	n := make([]string, nameLen, nameCap)
	fmt.Println(len(n), cap(n))*/

	/*num := 9
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v x %v = %v\t", i, j, i*j)
		}
		fmt.Println()
	}*/

	sumNum := 0
	for i := 1; i <= 100; i++ {
		sumNum += i
		if sumNum >= 20 {
			fmt.Println(i)
			break
		}
	}
}
