package main

import "fmt"

// 声明变量
/*var name string
var age int
var isOs string*/

// 批量声明 非全局变量被声明必须使用 推荐使用小驼峰式命名
var (
	name string //""       默认值
	age  int    //0
	isOk bool   //false
)

// 常量 声明后必须赋值 且不能修改
const through = 151.5481
const dimension = 139.1544

// 批量声明
const (
	success = 200
	fail    = 500
)

//	如果批量声明时如果某一行没有赋值，默认值为上一行所声明的值
const (
	n1 = 100
	n2
	n3
)

// iota声明常量 累加
const (
	key0 = iota //0
	key1        //1...
	key2
	key3
	key4
)

// 声明函数
func foo() (int, string) {
	return 10, "Q1mi"
}

func main() {
	//全局变量
	name = "测试"
	age = 22
	isOk = true

	//内部变量
	// 声明变量并赋值
	var inside string = "内部变量"
	// 类型推导 根据值判断该变量是什么类型
	var number = 120
	// 简短变量声明
	content := "内容"
	// 匿名变量 _ 用于占位
	x, _ := foo()

	fmt.Printf("name:%s", name) //%s占位符
	fmt.Println()               //输出换行
	fmt.Println(age)            //打印完自动换行
	fmt.Println(isOk)           //在终端中输出要打印的内容
	fmt.Println(inside)
	fmt.Println(number)
	fmt.Println(content)
	fmt.Print(x)
}
