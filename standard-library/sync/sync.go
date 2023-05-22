package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
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

	// sync.Cond实现了一个条件变量，用于等待一个或一组goroutines满足条件后唤起的场景，每个Cond都有一个关联的Locker L（通常是* Mutex或* RWMutex）
	// 另外一种写法创建：cond := &sync.Cond{L: new(sync.Mutex)}
	cond := sync.NewCond(new(sync.Mutex))
	// 阻塞当前的goroutine，等待唤起 必须遵守以下写法
	//    c.L.Lock() 上锁
	//    for !condition() { 不满足条件则阻塞goroutine执行
	//        c.Wait() 阻塞
	//    }
	//    ... make use of condition ... 利用条件（阻塞结束后业务逻辑...）
	//    c.L.Unlock() 解锁
	cond.Wait()
	// 唤起所有阻塞的goroutine
	cond.Broadcast()
	// 唤起第一个阻塞的goroutine
	cond.Signal()
	// 上锁
	cond.L.Lock()
	// 解锁
	cond.L.Unlock()

	// sync.Once是Golang package中使方法只执行一次的对象实现，作用与init函数类似。但也有所不同
	// init函数是在文件包首次被加载的时候执行，且只执行一次
	// sync.Onc是在代码运行中需要的时候执行，且只执行一次
	var once sync.Once
	// 自定义函数变量
	funcOne := func() { fmt.Println("函数只会被执行一次") }
	for i := 0; i < 3; i++ {
		// 执行函数 不管是在循环中还是多个协程中都只会被执行一次
		once.Do(funcOne)
	}
	// sync.Pool是对象缓存池，用来减少堆上内存的反复申请和释放的
	var pool = sync.Pool{
		// New方法定义默认的数据对象，如果没有设置则返回nil
		New: func() interface{} {
			return "default data"
		},
	}
	// Get方法从pool对象缓存池里获取任意一个对象 步骤：
	// 1.private不是空的，那就直接拿来用
	// 2.private是空的，那就先去本地的shared队列里面从头 pop 一个
	// 3.本地的shared也没有了，那getSlow去拿，其实就是去别的P的shared里面偷，偷不到就去victim幸存者里面找
	// 4.最后都没有，那就只能调用New方法创建一个了
	fmt.Println(pool.Get().(string)) // 因为返回值是interface{}，因此需要进行类型转换
	// Put方法将数据对象放进pool对象缓存池 步骤：
	// 1.private没有，就放在private
	// 2.private有了，那么就放到shared队列的头部
	pool.Put("put data")
	// 第一次GC垃圾回收会将pool对象缓存池内数据拷贝一份放进victim幸存者中,避免GC将其清空
	runtime.GC()
	time.Sleep(1 * time.Second)
	fmt.Println(pool.Get().(string)) // put data | default data 因为Get方法是获取任意一个对象，所以此处应有两种结果
	fmt.Println(pool.Get().(string)) // default data
	// 每个sync.Pool的生命周期为两次GC中间时段才有效 （最多存放两次GC）
	runtime.GC()
	time.Sleep(1 * time.Second)
	fmt.Println(pool.Get().(string)) // default data
}
