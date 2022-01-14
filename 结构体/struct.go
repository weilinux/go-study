package main

import "fmt"

//初始化
func init() {
	fmt.Println("初始化完成")
}

//结构体 值
type user struct {
	name string //值 “”
	age int //值 0
	hobby2 [2] string //值 [“”, “”]
	hobby []string //引用 nil slice切片 make开辟内存 append加长  值类型、len长度、cap容量
	aaa map[string]string //引用 nil map
}

type nums int

func main()  {
	//延迟执行
	defer func() {
		//捕捉异常
		if errorMsg := recover();errorMsg != nil {
			fmt.Printf("捕捉异常信息：%v\n", errorMsg)
		}
	}()

	fmt.Println("进入主函数")

	var user1 user
	fmt.Println(user1)
	user1.name = "小明"
	user1.age = 18
	user1.hobby2[0] = "游泳222"
	user1.hobby2[1] = "跑步333"
	user1.hobby = make([]string, 2)
	user1.hobby[0] = "游泳"
	user1.hobby[1] = "跑步"
	user1.aaa = make(map[string]string, 2)
	user1.aaa["a"] = "b"
	user1.aaa["c"] = "d"
	fmt.Println(user1)

	// 赋值
	user2 := user1
	fmt.Println(user2)

	user2.name = "小红"
	user2.age = 17
	user2.hobby2[0] = "游泳34653"
	user2.hobby2[1] = "跑步53633"
	user2.hobby[0] = "骑行"
	user2.hobby[1] = "游泳"
	user2.aaa["e"] = "f"
	fmt.Println(user1)
	fmt.Println(user2)

	// [0, 0, 0, 0, 0, 0] 6 10
	a222 := make([]int, 2, 5)
	fmt.Println(a222)

	num1 := nums(1)
	fmt.Println(num1)
	fmt.Printf("num1的数据类型是%T\n", num1)

	var num2 nums = 2
	fmt.Println(num2)
	fmt.Printf("num2的数据类型是%T\n", num2)

	panic("狗熊是傻逼")
}
