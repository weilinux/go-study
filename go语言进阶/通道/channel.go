package main

import (
	"fmt"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

func main() {
	//TODO 非缓冲通道 必须一存一读  如果通道为空则会阻塞等待插入数据后才会继续读取
	ch1 := make(chan int)
	fmt.Println(len(ch1), cap(ch1))

	//TODO 缓冲通道 创建时指定缓冲区大小，在缓冲区长度之内插入不会阻塞，但缓冲区已满或通道为空时读取会阻塞（等待插入数据）
	ch2 := make(chan string, 5)
	fmt.Println(len(ch2), cap(ch2))

	wg.Add(2)
	go sendData(ch2)
	go findData(ch2)

	//TODO 定向通道 规定通道的操作类型（读、写），应用于不同的业务场景
	ch3 := make(chan string)
	//ch4 := make(chan<- string)   只写
	//ch5 := make(<-chan string)   只读

	wg.Add(2)
	go writeData(ch3)
	go readData(ch3)

	wg.Wait()

	fmt.Println("程序执行完毕")
}

//缓冲通道写入
func sendData(ch chan string) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		ch <- "索引：" + strconv.Itoa(i)
		fmt.Println("缓冲 channel send:", i)
	}

	//使用完关闭通道
	close(ch)
}

//缓冲通过读取
func findData(ch chan string) {
	defer wg.Done()

	for {
		v, ok := <-ch
		if !ok {
			fmt.Println("缓冲 channel close:", ok)
			break
		}
		fmt.Println("缓冲 channel find:", v)
	}
}

//定向通道只写
func writeData(ch chan<- string) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		ch <- "索引：" + strconv.Itoa(i)
		fmt.Println("定向 channel write:", i)
	}

	//使用完关闭通道
	close(ch)
}

//定向通道只读
func readData(ch <-chan string) {
	defer wg.Done()

	for {
		v, ok := <-ch
		if !ok {
			fmt.Println("定向 channel close:", ok)
			break
		}
		fmt.Println("定向 channel read:", v)
	}
}
