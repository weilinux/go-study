package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	for i := 1; i <= 10; i++ {
		fmt.Println("func() 输出 " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func test2() {
	for i := 1; i <= 10; i++ {
		fmt.Println("func2() 输出 " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {
	//TODO  用goroutine（协程）输出函数

	go test() //开启一个协程

	go test2() //开启一个协程

	//协程开启后和主函数并行执行  不再由传统的从上到下执行
	for i := 1; i <= 10; i++ {
		fmt.Println("main() 输出 " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
