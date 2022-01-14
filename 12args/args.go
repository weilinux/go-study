package main

import (
	"fmt"
	"os"
)

func main() {
	//获取命令行参数
	fmt.Printf("获取命令行参数 类型：%T 值：%v\n", os.Args, os.Args)
}
