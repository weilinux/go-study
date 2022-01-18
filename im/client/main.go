package main

import (
	"fmt"
)

var (
	uid    int    //用户id
	pwd    string //密码
	option int    //接收用户的菜单选项
	loop   bool   //判断是否还继续显示菜单
)

//客户端主菜单
func home() {
	//TODO 客户端

	for {
		fmt.Println("--------------------欢迎登录多人聊天系统--------------------")
		fmt.Println("\t\t\t 1 登录聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择（1-3）")

		//获取用户输入
		_, err := fmt.Scanln(&option)
		if err != nil {
			panic(err)
		}

		switch option {
		case 1:
			fmt.Println("登录聊天室")
			loop = false
		case 2:
			fmt.Println("注册用户")
			loop = false
		case 3:
			fmt.Println("退出系统")
			loop = false
		default:
			fmt.Println("你的输入有误，请重新输入")
			loop = true
		}

		//如果选择正确则退出循环
		if loop == false {
			break
		}
	}
}

//用户选择结果
func result() {
	switch option {
	//登录
	case 1:
		fmt.Println("请输入用户id")
		_, err := fmt.Scanln(&uid)
		if err != nil {
			panic(err)
		}
		fmt.Println("请输入密码")
		_, err = fmt.Scanln(&pwd)
		if err != nil {
			panic(err)
		}

		//用户登录
		err = login(&uid, &pwd)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	//TODO 客户端主页

	//客户端主菜单
	home()

	//用户选择结果
	result()
}
