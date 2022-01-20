package main

import (
	"flag"
	"fmt"
)

func main() {
	//TODO 用flag包解析命令行参数

	//用户名
	var u string
	//密码
	var pwd string
	//主机
	var h string
	//端口
	var p int

	flag.StringVar(&u, "u", "", "用户名")
	flag.StringVar(&pwd, "pwd", "", "密码")
	flag.StringVar(&h, "h", "", "主机")
	flag.IntVar(&p, "p", 0, "端口")

	//解析命令行参数写入注册的flag里
	flag.Parse()

	fmt.Printf("用户名：%v 密码：%v 主机：%v 端口：%v \n", u, pwd, h, p)
}
