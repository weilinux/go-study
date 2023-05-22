package main

import (
	"fmt"
	"sync"
	"time"
)

//总票数
var ticket = 10

//创建锁头
var mutex sync.Mutex

//同步等待组对象
var wg sync.WaitGroup

func main() {
	//用4个goroutine模拟四个售票窗口

	//增加WaitGroup中的子goroutine计数值
	wg.Add(4)

	//模拟售票
	for i := 1; i <= 4; i++ {
		go saleTickets(i)
	}

	//阻塞调用此方法的goroutine，直到计数值为0
	wg.Wait()

	fmt.Println("售票完成")
}

//售票
func saleTickets(i int) {
	//当子goroutine任务完成，将计数值减1
	defer wg.Done()

	for {
		//上锁
		mutex.Lock()

		if ticket > 0 {
			time.Sleep(1 * time.Second)

			fmt.Printf("窗口：%d 售出：%d\n", i, ticket)

			//售出后减少票数
			ticket--
		} else {
			//无票也应解锁
			mutex.Unlock()

			fmt.Println("该票已售罄")

			break
		}

		//解锁
		mutex.Unlock()
	}
}
