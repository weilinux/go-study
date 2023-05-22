package main

import "fmt"

func f1() {
	fmt.Println("f1")
}

func f2() {
	//出现异常 提前释放资源
	defer func() {
		//尝试恢复 抑制错误 尽量少用 捕捉到的是抛出的异常信息
		if error := recover(); error != nil {
			fmt.Println(error)
			fmt.Printf("%T\n", error)
		}

		fmt.Println("关闭数据库连接")
	}()

	panic("异常错误")

	fmt.Println("f2")
}

func f3() {
	fmt.Println("f3")
}

func main() {
	f1()
	f2()
	f3()
}
