package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sort() {
	start := time.Now()

	arr := make([]int, 100)
	rand.Seed(start.Unix())
	for s := 0; s < 100; s++ {
		arr[s] = rand.Intn(100)
	}
	fmt.Printf("arr 排序前：%v\n", arr)

	//第一次循环遍历元素的个数（需要循环多少遍）
	for i := 0; i < len(arr)-1; i++ {
		//第二次循环遍历需要调整的个数（需要调整多少次，由于每次循环都能调整出当前循环体內最大的数所以第二遍循环需要遍历的个数每次都可以-1（-i））
		for j := 0; j < len(arr)-1-i; j++ {
			//如果当前元素大于索引下一位的元素则双方调换位置
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
