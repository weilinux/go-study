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

	//第一次循环遍历元素的个数（需要循环多少遍) 假设第一个元素默认是有序的，所以循环从第二个元素开始
	for i := 1; i < len(arr); i++ {
		//保存当前循环的元素
		current := arr[i]
		//保存当前循环初始需要对比的索引（i-1为从当前循环元素的前一位开始）
		j := i - 1
		//由于每次对比都是从当前元素前一位开始向左位移，所以第二次循环只需要确保下标移动到0（数组首位元素）并且第二次循环的元素大于当前循环元素
		for j >= 0 && arr[j] > current {
			//如果第二次循环的元素大于当前循环元素，则将第二次循环索引的元素赋值给索引+1的元素（后一位）
			arr[j+1] = arr[j]
			//循环递减 （索引位由大到小）
			j--
		}
		//如果索引并且发生更改则无需赋值（当前循环的元素大于第二次循环对比的第一个元素；因为条件不成立也会-1，所以该位置判断需要+1才是第一个元素的索引位）
		if j+1 != i {
			//条件成立则将当前循环的元素保存在条件不成立的位置+1处（第二次循环不成立后的索引位）
			arr[j+1] = current
		}
	}

	fmt.Printf("arr 排序后：%v\n", arr)

	elapsed := time.Since(start)
	fmt.Println("该函数执行完成耗时：", elapsed)
}

func main() {
	sort()
}
