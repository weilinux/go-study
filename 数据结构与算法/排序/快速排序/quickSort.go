package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
 * desc：快速排序（递归）
 * param：arr 需要排序的数组切片
 * param：start 数组初始位置
 * param：end 数组末尾位置
 */
func sort(arr []int, start, end int) {
	//如果初始位置小于末尾位置则一直对比并调用位置
	if start < end {
		//先将首位和末位赋值给两个变量（后续需递增或递减来进行值对比）
		left, right := start, end
		//数组中间位置（左右两边的数依次和该数对比）
		pivot := arr[(left+right)/2]

		//如果左索引小于等于右索引则一直循环（大于则排序完成）
		for left <= right {
			//如果数组左索引位置的值小于中间位置的值则左索引递增
			for arr[left] < pivot {
				left++
			}

			//如果数组右索引位置的值大于中间位置的值则右索引递减
			for arr[right] > pivot {
				right--
			}

			//如果左索引小于等于右索引则交换两边的值并让索引向中心进一位（大于则排序完成）
			if left <= right {
				arr[left], arr[right] = arr[right], arr[left]
				left++
				right--
			}
		}

		//如果数组初始位置小于右索引位置则递归进行排序
		if start < right {
			sort(arr, start, right)
		}

		//如果数组末尾位置大于左索引位置则递归进行排序
		if end > left {
			sort(arr, left, end)
		}
	}
}

func main() {
	start := time.Now()

	arr := make([]int, 100)
	rand.Seed(start.Unix())
	for s := 0; s < 100; s++ {
		arr[s] = rand.Intn(100)
	}
	fmt.Printf("arr 排序前：%v\n", arr)

	arr = []int{3, 9, 1, 7, 2, 5, 8, 4}

	sort(arr, 0, len(arr)-1)

	fmt.Printf("arr 排序后：%v\n", arr)

	elapsed := time.Since(start)
	fmt.Println("该函数执行完成耗时：", elapsed)
}
