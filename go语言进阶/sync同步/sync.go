package main

import (
	"fmt"
	"sync"
)

func main() {
	//TODO sync包详解

	// 协程等待组 等待组等待goroutine集合完成
	wg := sync.WaitGroup{}
	// 添加协程执行数量
	wg.Add(1)
	// 协程执行完毕将WaitGroup计数器递减1
	wg.Done()
	// 协程未执行完成则阻塞主函数结束运行
	wg.Wait()

	// 创建并发安全的map
	// 1、以空间换效率，通过read和dirty两个map来提高读取效率
	// 2、优先从read map中读取(无锁)，否则再从dirty map中读取(加锁)
	// 3、动态调整，当misses次数过多时，将dirty map提升为read map
	// 4、延迟删除，删除只是为value打一个标记，在dirty map提升时才执行真正的删除
	var smp sync.Map
	// 数据写入
	smp.Store("name", "jack")
	smp.Store("age", 18)
	// 数据读取
	name, ok := smp.Load("name")
	fmt.Println("msp 键name的值为：", name, "是否读取成功：", ok)
	// 遍历 如果f返回false，range将停止迭代
	smp.Range(func(key, value interface{}) bool {
		fmt.Println("msp 键：", key)
		fmt.Println("msp 值：", value)
		return true
	})
	// 删除
	smp.Delete("name")
	// 读取或者写入 存在就读取，不存在则写入
	actual, ok := smp.LoadOrStore("name", "tom2")
	fmt.Println("键name的值为：", actual, "是否写入成功（true：新增 false：修改）：", ok)
	// 读取并删除
	value, ok := smp.LoadAndDelete("age")
	fmt.Println("删除前 键age的值为：", value, "是否删除成功：", ok)

	// 创建互斥锁
	var mutex sync.Mutex
	// 上锁
	mutex.Lock()
	// 解锁
	mutex.Unlock()

	// 创建读写锁
	var rwMutex sync.RWMutex
	// 写操作上锁
	rwMutex.Lock()
	// 读操作上锁
	rwMutex.RLock()
	// 写操作解锁
	rwMutex.Unlock()
	// 读操作解锁
	rwMutex.RUnlock()
	// 返回一个实现了Lock()和Unlock()方法的Locker接口
	locker := rwMutex.RLocker()
	// 上锁(读操作)
	locker.Lock()
	// 解锁(读操作)
	locker.Unlock()

	// sync.Cond实现了一个条件变量，用于等待一个或一组goroutines满足条件后唤醒的场景
	var cond sync.Cond
	// 挂起当前执行的goroutine （自动调用cond.L.Lock()上锁方法）
	cond.Wait()
	// 唤醒所有被挂起的goroutine （自动调用cond.L.Unlock()解锁方法）
	cond.Broadcast()
	// 唤醒第一个被挂起的goroutine
	cond.Signal()

	/*var once sync.Once
	var pool sync.Pool*/
}
