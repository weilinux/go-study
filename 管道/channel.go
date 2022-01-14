package main

import "fmt"

func main() {
	//TODO channel管道  用于线程安全数据传输和协程一起使用  当没有数据时取值则会报错，数据放满后再放值也会报错

	//创建长度为10的管道
	intChan := make(chan int, 10)

	//向管道插入数据
	intChan <- 111
	intChan <- 222

	//获取管道长度、容量
	fmt.Printf("管道长度：%v 管道容量：%v\n", len(intChan), cap(intChan))

	//获取管道数据
	int1 := <-intChan
	fmt.Printf("管道数据1：%v\n", int1)
	fmt.Printf("管道长度：%v 管道容量：%v\n", len(intChan), cap(intChan))

	int2 := <-intChan
	fmt.Printf("管道数据1：%v\n", int2)
	fmt.Printf("管道长度：%v 管道容量：%v\n", len(intChan), cap(intChan))

	//关闭管道 管道关闭后仍可以向该管道里取数据，但无法再插入数据
	//close(intChan)
}
