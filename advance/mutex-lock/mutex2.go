package main

import (
	"fmt"
	"sync"
)

var (
	mutex2 sync.Mutex
	// 多个协程并发访问单一资源时，避免出现数据访问竞争，导致逻辑错误
	balance int
)

func deposit(value int, wg *sync.WaitGroup) {
	mutex2.Lock()
	fmt.Printf("Depositing %d to account with balance %d", value, balance)
	balance += value
	mutex2.Unlock()
	wg.Done()
}

func withdraw(value int, wg *sync.WaitGroup) {
	mutex2.Lock()
	fmt.Printf("withdraw %d to account with balance %d", value, balance)
	balance -= value
	mutex2.Unlock()
	wg.Done()

}

func main() {
	var wg sync.WaitGroup
	balance = 1000
	wg.Add(2)
	go withdraw(700, &wg)
	go deposit(500, &wg)
	wg.Wait()
	fmt.Println(balance)
}
