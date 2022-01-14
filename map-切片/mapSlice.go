package main

import (
	"fmt"
	"strings"
)

func main() {
	// TODO 切片中含有map
	// 初始化元素类型为map的切片
	ms1 := make([]map[int]string, 10, 10)

	// 初始化切片里的map
	ms1[0] = make(map[int]string, 1)

	ms1[0][0] = "slice>map"
	fmt.Println(ms1)

	// TODO map中含有切片
	// 初始化元素类型为切片的map
	sm1 := make(map[int][]string, 10)

	// 初始化map中的切片
	sm1[0] = make([]string, 10) //此处可不用填写 赋值的同时默认初始化 单赋值之前初始化更规范 并指定切片大小 以防内存溢出

	sm1[0] = []string{"北京", "上海", "广州", "深圳"}
	fmt.Println(sm1)

	// TODO 统计每个不同的单词在字符串中出现的次数
	s1 := "how do you do"
	fmt.Println(s1)

	// 根据“ ”（空格）对字符串进行分割
	sp1 := strings.Split(s1, " ")
	fmt.Println(sp1)

	// 循环切片对相同值进行累加
	m1 := make(map[string]int, 3)
	for _, value := range sp1 {
		m1[value] += 1
	}
	fmt.Println(m1)
}
