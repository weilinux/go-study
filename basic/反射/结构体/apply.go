package main

import (
	"fmt"
	"reflect"
)

// Monster 结构体
type Monster struct {
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Score float32 `json:"score"`
	Sex   string  `json:"sex"`
}

// Print 打印结构体信息
func (m Monster) Print() {
	fmt.Println("----start---")
	fmt.Println(m)
	fmt.Println("-----end----")
}

// GetSum 计算两个数的和
func (m Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

// Set 给Monster赋值
func (m Monster) Set(name string, age int, score float32, sex string) {
	m.Name = name
	m.Age = age
	m.Score = score
	m.Sex = sex
}

// TestStruct 反射
func TestStruct(i interface{}) {
	//实例反射结构体类型
	iType := reflect.TypeOf(i)

	//实例反射结构体值
	iValue := reflect.ValueOf(i)

	//获取反射值类别
	iKind := iValue.Kind()

	//如果不是结构体就打断程序执行
	if iKind != reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	//获取结构体字段数量
	iNumField := iValue.NumField()
	fmt.Printf("struct has %d fields\n", iNumField)

	//遍历结构体字段
	for i := 0; i < iNumField; i++ {
		//获取对应下标的值
		fmt.Printf("field %d: 值为：%v\n", i, iValue.Field(i))

		//获取对应下标的tag标签：json
		iTag := iType.Field(i).Tag.Get("json")
		if iTag != "" {
			fmt.Printf("field %d: tag为=%v\n", i, iTag)
		}
	}

	//获取结构体方法数量
	iNumMethod := iValue.NumMethod()
	fmt.Printf("struct has %d methods\n", iNumMethod)

	//Method：调用结构体的方法 call：方法参数，数据格式为[]Value切片
	//方法的默认排序是按照函数名（Ascii码的大小进行sort）
	iValue.Method(1).Call(nil)

	//调用结构体的第一个方法：method(0)
	params := make([]reflect.Value, 2)
	params[0] = reflect.ValueOf(10)
	params[1] = reflect.ValueOf(40)
	res := iValue.Method(0).Call(params)
	fmt.Println("res=", res[0].Int())

	//获取结构体方法名对应地址
	iMethodName := iValue.MethodByName("Set")
	fmt.Println("结构体方法地址为：", iMethodName)

	//调用此方法
	params2 := make([]reflect.Value, 4)
	params2[0] = reflect.ValueOf("牛魔王")
	params2[1] = reflect.ValueOf(1500)
	params2[2] = reflect.ValueOf(float32(60.5))
	params2[3] = reflect.ValueOf("公")
	res2 := iMethodName.Call(params2)
	fmt.Println("res2=", res2)
}

func main() {
	//实例结构体
	m := Monster{
		Name:  "黄鼠狼",
		Age:   500,
		Score: 99.5,
		Sex:   "公",
	}

	//反射结构体进行类型推断和赋值
	TestStruct(m)
}
