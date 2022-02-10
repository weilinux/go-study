package main

import (
	"fmt"
	"sort"
)

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 3, 2}
	fmt.Println(s1)
	fmt.Printf("%T", s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
	fmt.Println(s1[5:])
	fmt.Println()

	s2 := []string{"啊", "哦", "额", "咦", "唔", "喻"}
	fmt.Println(s2)
	fmt.Printf("%T", s2)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))
	fmt.Println(s2[len(s2)-1])
	fmt.Println(s2[0:4])
	fmt.Println(s2[:4])
	fmt.Println()

	// make()创建切片
	s3 := []string{"月", "上", "柳", "梢", "头", "人"}
	fmt.Printf("len => %v cap => %v val => %v\n", len(s3), cap(s3), s3)

	// append()追加元素 数组长度小于等于1024时 超出容量会自动扩充1倍  大于1024时 超出容量会自动扩充1.25倍 超出容量后才会返回新的数组
	s3 = append(s3, "约")
	fmt.Printf("len => %v cap => %v val => %v\n", len(s3), cap(s3), s3)
	s3 = append(s3, "黄")
	fmt.Printf("len => %v cap => %v val => %v\n", len(s3), cap(s3), s3)
	s4 := append(s3, "昏")
	fmt.Printf("len => %v cap => %v val => %v\n", len(s3), cap(s3), s3)
	fmt.Printf("len => %v cap => %v val => %v\n", len(s4), cap(s4), s4)
	s3 = append(s3, s2...) // ... 解析赋值
	fmt.Printf("len => %v cap => %v val => %v\n", len(s3), cap(s3), s3)

	// 浅拷贝 指针指向同一数组
	s5 := s3
	// copy() 近似于覆盖 有值则覆盖对应下标的值 且拷贝的切片数据为初始数组 通过append()方法追加的元素不会拷贝 重新分配内存 创建新数组储存
	s6 := make([]string, 6, 6)
	copy(s6, s3)
	fmt.Println(s5)
	fmt.Println(s6)

	// 删除切片元素
	sep := 4
	s7 := append(s3[:sep], s3[sep+1:]...)
	fmt.Println(s7)

	// sort.Ints()对数值类型的切片进行排序
	sort.Ints(s1)
	fmt.Println(s1)
}
