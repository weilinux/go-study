package main

import "fmt"

// TODO 函数类型
// 无参数无让返回值格式
func sum1() {
	fmt.Println("无参数无返回值函数")
}

// 无返回值格式
func sum2(x int, y int) {
	fmt.Println(x + y)
}

// 无参数格式
func sum3() int {
	x := 10
	y := 13

	return x + y
}

// 返回定义返回数据类型格式
func sum4() int {
	return 1
}

// 返回命名返回变量名和返回数据类型格式 定义的变量名为函数中声明的变量 可以在函数体内使用
func sum5() (ret int) {
	ret = 2

	// 如果为此格式的函数类型 只需写上return即可完成数据返回
	return
}

// 多返回参数格式 如果数据类型相同可简写
func sum6() (ret1, ret2 int, ret3 string) {
	return 3, 4, "多返回参数格式"
}

// 可变长参数格式 必须写在最后一个
func sum7(class int, student ...string) {
	fmt.Println(class)
	fmt.Println(student)
}

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
