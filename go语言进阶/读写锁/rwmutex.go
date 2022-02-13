package main

import (
	"fmt"
	"sync"
)

//创建锁头
var rwMutex sync.RWMutex

//同步等待组对象
var wg sync.WaitGroup

func main() {
	//增加WaitGroup中的子goroutine计数值
	wg.Add(8)

	//模拟读写操作
	for i := 1; i <= 4; i++ {
		go read(i)
		go write(i)
	}

	//阻塞调用此方法的goroutine，直到计数值为0
	wg.Wait()

	fmt.Println("程序执行完毕")
}

//读操作
func read(i int) {
	//当子goroutine任务完成，将计数值减1
	defer wg.Done()

	fmt.Printf("用户：%d 读操作start\n", i)

	//读操作上锁
	rwMutex.RLock()

	fmt.Printf("用户：%d 正在读取数据\n", i)

	//读操作解锁
	rwMutex.RUnlock()

	fmt.Printf("用户：%d 读操作end\n", i)
}

//写操作
func write(i int) {
	//当子goroutine任务完成，将计数值减1
	defer wg.Done()

	fmt.Printf("用户：%d 写操作start\n", i)

	//写操作上锁
	rwMutex.Lock()

	fmt.Printf("用户：%d 正在写入数据\n", i)

	//写操作解锁
	rwMutex.Unlock()

	fmt.Printf("用户：%d 写操作end\n", i)
}
