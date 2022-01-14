package main

import "fmt"

var test string

func main() {

	for {
		if test != "" {
			fmt.Println("再次输入试试")
		} else {
			fmt.Println("请输入栗木村几个鬼的名字")
		}
		fmt.Scan(&test)

		switch test {
		case "威子":
			fmt.Println("内真滴就是个卵鬼")
			fmt.Println()
		case "狗熊":
			fmt.Println("前进村最虚伪滴就是他")
			fmt.Println()
		case "嗨子":
			fmt.Println("著名南派相声演员向科的大徒弟，取名廖天嗨，代表作：《探澧水河》、《澧南舞厅》")
			fmt.Println()
		case "暴风":
			fmt.Println("内个儿搞不好哒滴")
			fmt.Println()
		case "康师傅":
			fmt.Println("来分保险，不加方便面")
			fmt.Println()
		case "重机子":
			fmt.Println("重型机器已启动...\"啊\" 三米高~~~")
			fmt.Println()
		case "科子":
			fmt.Println("春晚舞台：科：我和你 嗨：心连心 合唱：同住前进村")
			fmt.Println()
		}
	}
}
