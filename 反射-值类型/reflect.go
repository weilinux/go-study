package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

func getType(i interface{}) {
	//获取变量类型
	iType := reflect.TypeOf(i)
	fmt.Printf("获取变量类型：%v\n", iType)
	fmt.Printf("获取变量名称：%v\n", iType.Name())

	//获取变量值
	iValue := reflect.ValueOf(i)
	fmt.Printf("获取变量值：%v\n", iValue)

	//获取变量类型
	vType := iValue.Type()
	fmt.Printf("获取变量类型：%v\n", vType)

	//获取变量类别（变量本身类型的常量，例如：int、string、slice、map...）
	vKind := iValue.Kind()
	fmt.Printf("获取变量类别：%v\n", vKind)

	//修改对应的变量值
	iValue.Elem().SetInt(20)
	fmt.Printf("修改对应的变量值：%v\n", iValue)

	//将该地址转成接口类型
	i2 := iValue.Interface()
	fmt.Printf("将该地址转成接口类型：%v\n", i2)

	//将此接口类型地址通过类型断言转换成对应类型的变量
	user2 := i2.(*int)
	fmt.Printf("将此接口类型地址通过类型断言转换成对应类型的变量：%v\n", *user2)
}

func main() {
	//TODO 反射

	/*user1 := User{
		Name: "小明",
		Age:  18,
	}

	getType(user1)*/

	i1 := 10
	getType(&i1)
	fmt.Printf("main() 打印变量值：%v\n", i1)

}
