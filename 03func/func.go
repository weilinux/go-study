package main

import "fmt"

/*func swap(a, b *int) {
	*a, *b = *b, *a
}*/

func addNum(num int) func(int, int) int {
	number := 0

	return func(incr, step int) int {
		number += incr*step + num
		return number
	}
}

func main() {
	/*a, b := 1, 2
	fmt.Printf("a 初始值：%v\n", a)
	fmt.Printf("b 初始值：%v\n", b)

	swap(&a, &b)
	fmt.Printf("a 调整后值：%v\n", a)
	fmt.Printf("b 调整后值：%v\n", b)*/

	a := addNum(10)
	fmt.Println(a(1, 2))
	fmt.Println(a(2, 3))
	fmt.Println(a(3, 4))
}
