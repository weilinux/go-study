package main

import (
	"fmt"
	"strings"
)

func main() {
	// 字符串 \需要转移
	s1 := "https://www.baidu.com/search"
	fmt.Println(s1)

	// 格式化字符串 不需要转移
	s2 := `
		月上柳梢头
		人约黄昏后
	`
	fmt.Println(s2)

	// 获取字符串长度
	fmt.Println(len(s1))

	// 字符串拼接
	name := "小明"
	action := "爱学习"
	s3 := name + action
	fmt.Println(s3)
	//另外一种拼接方式
	user := fmt.Sprintf("%s%s", name, action)
	fmt.Println(user)

	// 字符串分割
	urlSpint := strings.Split(s1, "/")
	fmt.Println(urlSpint)

	// 判断该字符串是否存在与指定的字符串变量中 存在返回true 否则返回false
	fmt.Println(strings.Contains(s2, "月"))
	fmt.Println(strings.Contains(s2, "前"))

	// 判断该字符串是否是指定的字符串变量的前缀:后缀
	fmt.Println(strings.HasPrefix(s2, ""))
	fmt.Println(strings.HasSuffix(s2, "人"))

	// 判断该字符串在指定的字符串变量中的位置 存在则返回index位置，默认从0开始 不存在则返回-1
	fmt.Println(strings.Index(s1, "h"))
	fmt.Println(strings.LastIndex(s1, "f"))

	// 数组拼接成字符串
	fmt.Println(strings.Join(urlSpint, "\\"))
}
