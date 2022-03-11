package main

import (
	"fmt"
	"runtime"
)

func main() {
	//TODO runtime包详解

	// 获取GOROOT环境变量根目录
	fmt.Println("GOROOT环境变量根目录：", runtime.GOROOT())

	// 获取Go语言版本号
	fmt.Println("Go语言版本号：", runtime.Version())

	// 设置最大可同时执行的最大CPU核数 默认为当前系统CPU最大核数，当参数n小于1时使用默认值
	i := runtime.GOMAXPROCS(1)
	fmt.Println("上一个设置的CPU核数为：", i)

	// 返回当前系统的CPU核数
	fmt.Println("当前系统的CPU核数为：", runtime.NumCPU())

	// 返回当前进程进行的cgo调用数 (当前进程调用c方法的次数)
	fmt.Println("当前进程进行的cgo调用数", runtime.NumCgoCall())

	// 返回当前存在的goroutine数
	fmt.Println("当前存在的goroutine数", runtime.NumGoroutine())

	// 设置CPU分析速率设置为每秒hz采样数 如果hz<=0，SetCpurpFileRate将关闭评测
	runtime.SetCPUProfileRate(100)

	// GC运行垃圾收集并阻塞程序运行，直到垃圾收集完成
	//runtime.GC()

	// 给变量绑定方法,当垃圾回收的时候进行监听 第一次不会销毁并触发执行该变量绑定的方法，第二次则会销毁 （一般x为变量地址，y为该变量绑定的方法）
	//runtime.SetFinalizer()

	// 让当前执行的goroutine让出CPU时间片，重新等待安排任务 (让其他go协程优先执行,等其他协程执行完后,在执行当前的协程)
	//runtime.Gosched()

	// 终止当前goroutine运行，终止调用该goroutine的运行之前会先执行该goroutine中还没有执行的defer语句
	//runtime.Goexit()

	// 查看内存申请和分配统计信息
	memStats := &runtime.MemStats{}
	runtime.ReadMemStats(memStats)
	/*type MemStats struct {
		// 一般统计
		Alloc      uint64 // 已申请且仍在使用的字节数
		TotalAlloc uint64 // 已申请的总字节数（已释放的部分也算在内）
		Sys        uint64 // 从系统中获取的字节数（下面XxxSys之和）
		Lookups    uint64 // 指针查找的次数
		Mallocs    uint64 // 申请内存的次数
		Frees      uint64 // 释放内存的次数
		// 主分配堆统计
		HeapAlloc    uint64 // 已申请且仍在使用的字节数
		HeapSys      uint64 // 从系统中获取的字节数
		HeapIdle     uint64 // 闲置span中的字节数
		HeapInuse    uint64 // 非闲置span中的字节数
		HeapReleased uint64 // 释放到系统的字节数
		HeapObjects  uint64 // 已分配对象的总个数
		// L低层次、大小固定的结构体分配器统计，Inuse为正在使用的字节数，Sys为从系统获取的字节数
		StackInuse  uint64 // 引导程序的堆栈
		StackSys    uint64
		MSpanInuse  uint64 // mspan结构体
		MSpanSys    uint64
		MCacheInuse uint64 // mcache结构体
		MCacheSys   uint64
		BuckHashSys uint64 // profile桶散列表
		GCSys       uint64 // GC元数据
		OtherSys    uint64 // 其他系统申请
		// 垃圾收集器统计
		NextGC       uint64 // 会在HeapAlloc字段到达该值（字节数）时运行下次GC
		LastGC       uint64 // 上次运行的绝对时间（纳秒）
		PauseTotalNs uint64
		PauseNs      [256]uint64 // 近期GC暂停时间的循环缓冲，最近一次在[(NumGC+255)%256]
		NumGC        uint32
		EnableGC     bool
		DebugGC      bool
		// 每次申请的字节数的统计，61是C代码中的尺寸分级数
		BySize [61]struct {
			Size    uint32 // 此模式中对象的最大字节大小
			Mallocs uint64 // 堆对象的累积字节数
			Frees   uint64 // 释放的堆对象的累积字节数
		}
	}*/

	// 查看goroutine堆栈信息
	m := &runtime.MemProfileRecord{}
	// 返回正在使用的字节数
	fmt.Println("正在使用的字节数", m.InUseBytes())
	// 返回正在使用的字节数
	fmt.Println("正在使用的对象数", m.InUseObjects())
	// 获取调用堆栈列表
	fmt.Println("关联至此记录的调用栈踪迹", m.Stack())

	//获取程序调用go协程的堆栈踪迹历史
	buf := make([]byte, 1024)
	runtime.Stack(buf, true)
	fmt.Println("堆栈信息：", string(buf))

	// 获取当前函数或者上层函数的标识号、文件名、调用方法在当前文件中的行号 升序，0表示调用者的调用者
	pc, file, line, ok := runtime.Caller(0)
	fmt.Println("标识号：", pc, " 文件名：", file, " 调用方法", line, " 是否成功：", ok)

	// 获取与当前堆栈记录相关链的调用栈踪迹 0表示调用者本身的帧 1表示调用者中的调用者
	pcs := make([]uintptr, 10)
	i = runtime.Callers(0, pcs)
	fmt.Println(pcs[:i])

	// 获取一个标识调用栈标识符pc对应的调用栈
	funcPc := runtime.FuncForPC(pc)
	fmt.Println("调用栈标识符pc对应的调用栈：", funcPc)
	// 获取调用栈所调用的函数的名字
	fmt.Println("调用栈所调用的函数的名字：", funcPc.Name())
	// 获取调用栈所调用的函数的所在的源文件名和行号
	file, line = funcPc.FileLine(pc)
	fmt.Println("调用栈所调用的函数的所在的源文件名和行号：", file, line)
	// 获取调用栈的调用栈标识符
	fmt.Println("调用栈的调用栈标识符：", funcPc.Entry())

	// 获取活跃的goroutine堆栈配置文件中的记录数
	n, ok := runtime.GoroutineProfile(make([]runtime.StackRecord, 10))
	fmt.Println("活跃的goroutine堆栈配置文件中的记录数：", n, "是否成功：", ok)

	// 将调用的go协程绑定到当前所在的操作系统线程，其它go协程不能进入该线程
	go func() {
		// 解除go协程与操作系统线程的绑定关系，否则它将总是在该线程中执行 (若调用的go程未调用LockOSThread，UnlockOSThread不做操作)
		defer runtime.UnlockOSThread()

		// 绑定
		runtime.LockOSThread()

		fmt.Println("goroutine绑定线程中...")
	}()

	// 获取线程创建profile中的记录个数
	n, ok = runtime.ThreadCreateProfile(make([]runtime.StackRecord, 10))
	fmt.Println("线程创建profile中的记录个数：", n, "是否成功：", ok)

	// 控制阻塞profile记录go协程阻塞事件的采样率 要在profile中包括每一个阻塞事件，需传入rate=1；要完全关闭阻塞profile的记录，需传入rate<=0
	runtime.SetBlockProfileRate(0)

	// 返回当前阻塞profile中的记录个数
	n, ok = runtime.BlockProfile(make([]runtime.BlockProfileRecord, 10))
	fmt.Println("当前阻塞profile中的记录个数：", n, "是否成功：", ok)
}
