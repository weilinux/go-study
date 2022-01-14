package main

import (
	"fmt"
	"sync"
)

func main() {
	//TODO WaitGroup对象内部有个计时器， 最初从0 开始， 他有3个方法 Add() , Done(), Wait()用来控制计数器的数量。 Add(n) 把计数器设置成n, Done() 每次把计数器-1， wait() 会阻塞代码的运行， 直到计数器的值减为0

	//将对象实例化
	wg := sync.WaitGroup{}

	//添加协程执行数量
	wg.Add(10)

	/*for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)

			//执行完成后减少此程序
			wg.Done()
		}(i)
	}*/

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("狗熊是帅锅  当前循环下标：", i)

			//执行完成后减少此程序
			wg.Done()
		}(i)
	}

	//协程未执行完成则阻塞主函数结束运行
	wg.Wait()
}
