package main

import (
	"fmt"
	"runtime"
)

func main() {
	//TODO 查看系统CPU个数，并分配go程序执行个数

	//系统CPU个数
	numCPU := runtime.NumCPU()
	fmt.Println(numCPU)

	//设置程序运行的cpu个数
	result := runtime.GOMAXPROCS(numCPU - 1)
	fmt.Printf("程序运行的CPU个数为：%v\n", result)
}
