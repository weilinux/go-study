package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	// 初始化map 长度为10
	m1 := make(map[string]string, 10)

	// 赋值
	m1["name"] = "小明"
	m1["city"] = "长沙"
	fmt.Println(m1)
	fmt.Println(m1["name"])
	fmt.Println(len(m1))

	// 不存在的键则返回对应声明类型的初始值 "" 0 false
	fmt.Println(m1["age"])

	// 判断此键是否存在与map变量中 存在返回true 否则返回false
	value, ok := m1["age"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("不存在此键")
	}

	// 遍历map
	for k, v := range m1 {
		fmt.Printf("%v => %v\n", k, v)
	}

	// 删除指定的键
	delete(m1, "city")
	fmt.Println(m1)

	// 初始化随机数种子
	rand.Seed(time.Now().Unix())
	// 初始化长度为200的map变量
	scoreMap := make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		// 生成以stu开头的字符串
		key := fmt.Sprintf("stu%02d", i)
		// 生成0~99区间的随机数
		value := rand.Intn(100)
		// 拼接key-value
		scoreMap[key] = value
	}
	fmt.Println(scoreMap)

	// 提取出所有的key
	scoreKeys := make([]string, 0, 200)
	for key := range scoreMap {
		scoreKeys = append(scoreKeys, key)
	}
	fmt.Println(scoreKeys)

	// 将切片进行排序
	sort.Strings(scoreKeys)
	fmt.Println(scoreKeys)

	// 按照排序后的key遍历map
	for _, valueKey := range scoreKeys {
		fmt.Println(valueKey, scoreMap[valueKey])
	}
}
