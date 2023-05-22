package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("程序初始时间：", time.Now().Format("2006-01-02 15:04:05"))

	//TODO timer定时器，固定时间后会执行一次
	//设置定时器为3秒
	timer := time.NewTimer(2 * time.Second)
	wg.Add(1)
	go func(t *time.Timer) {
		defer wg.Done()

		//停止计时器 如果不重置则会一直阻塞
		t.Stop()

		//重置定时器时间 还在等待中会返回真；如果t已经到期或者停止了会返回假。
		t.Reset(5 * time.Second)

		<-t.C
		fmt.Println("get timer", time.Now().Format("2006-01-02 15:04:05"))
	}(timer)

	//TODO ticker只要定义完成，从此刻开始计时，不需要任何其他的操作，每隔固定时间都会
	wg.Add(1)
	//设置定时器为2秒
	ticker := time.NewTicker(2 * time.Second)
	go func(t *time.Ticker) {
		defer wg.Done()

		//设置执行十次后中断
		i := 0
		for {
			<-t.C
			fmt.Println("get ticker", time.Now().Format("2006-01-02 15:04:05"))

			i++
			if i == 10 {
				break
			}
		}
	}(ticker)

	//TODO sleep阻塞函数运行 timer方法的封装，直接返回时间类型的通道
	wg.Add(1)
	go func() {
		defer wg.Done()

		fmt.Println("阻塞前:", time.Now().Format("2006-01-02 15:04:05"))

		time.Sleep(3 * time.Second)

		fmt.Println("阻塞后:", time.Now().Format("2006-01-02 15:04:05"))
	}()

	//TODO after函数实现延迟功能
	wg.Add(1)
	//设置定时器为2秒
	after := time.After(2 * time.Second)
	go func(t <-chan time.Time) {
		defer wg.Done()

		<-t
		fmt.Println("get after", time.Now().Format("2006-01-02 15:04:05"))
	}(after)

	//TODO afterFunc函数实现延迟传入函数来执行
	//设置定时器为3秒
	time.AfterFunc(3*time.Second, func() {
		fmt.Println("get afterFunc", time.Now().Format("2006-01-02 15:04:05"))
	})

	wg.Wait()
}
