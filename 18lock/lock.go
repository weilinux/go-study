package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

var (
	sum  int
	lock sync.Mutex
)

//获取协程id
func getGoroutineID() uint64 {
	b := make([]byte, 64)
	runtime.Stack(b, false)
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func sumNum(n int) {
	//fmt.Println(getGoroutineID())

	//加锁 保证数据无协程执行的冲突
	lock.Lock()
	sum += n
	//解锁
	lock.Unlock()
}

func main() {
	num := 100
	for i := 0; i <= num; i++ {
		go sumNum(num)
	}

	defer func(i int) {
		fmt.Println(i)
	}(sum)
}
