package main

import (
	model "../model"
	service "../service"
	"fmt"
)

var aaa int

func init() {
	aaa = 1
}

func main() {
	fmt.Println(aaa)

	member1 := model.AddMember("tom", 18, true)
	member1.SetName("新名字")
	service.AgeInc(member1)

	member2 := model.AddMember("jerry", 20, false)
	member2.SetName("新新名字")
	service.AgeInc(member2)

	fmt.Println(*member1)
	fmt.Println(*member2)

	member3 := model.Member2{}
	fmt.Printf("结构体的初始值： %v\n", member3)

	//空接口
	m1 := make(map[string]interface{})
	m1["a"] = 1
	m1["b"] = "b"
	m1["c"] = true
	m1["d"] = 1.2
	m1["e"] = []int{1, 2, 3, 4, 5}
	fmt.Println(m1)

	//类型断言
	m2, ok := m1["a"].(int)
	if !ok {
		fmt.Println("类型不正确")
	}
	fmt.Println(m2)

	m3, ok := m1["e"].([]int)
	if !ok {
		fmt.Println("类型不正确")
	}
	fmt.Println(m3)
}
