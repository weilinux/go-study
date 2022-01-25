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
		//保存当前循环最小值的索引位（初始值为i）
		min := i
		//第二次循环遍历需要调整的个数（需要调整多少次，由于每次循环都能调整出当前循环体內最小的数所以第二次循环只需要从最小值索引位+1的数值遍历数组（i+i为数组头部已确认，无需再改动）））
		for j := i + 1; j < len(arr); j++ {
			//如果保存的最小值比当前值大则双方交换位置
			if arr[min] > arr[j] {
				min = j
			}
		}
		//如果最小值索引发生改变已交换位置则双方调换位置
		if min != i {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}

	fmt.Printf("arr 排序后：%v\n", arr)

	elapsed := time.Since(start)
	fmt.Println("该函数执行完成耗时：", elapsed)
}

func main() {
	sort()
}
