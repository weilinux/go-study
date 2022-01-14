package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sort() {
	start := time.Now()

	arr := make([]int, 10000)
	rand.Seed(start.Unix())
	for s := 0; s < 10000; s++ {
		arr[s] = rand.Intn(10000)
	}
	fmt.Printf("arr 排序前：%v\n", arr)

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	fmt.Printf("arr 排序后：%v\n", arr)

	elapsed := time.Since(start)
	fmt.Println("该函数执行完成耗时：", elapsed)
}

func main() {
	sort()
}
