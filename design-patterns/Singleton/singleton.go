package main

import (
	"fmt"
	"reflect"
	"sync"
)

// Singleton 业务对象
type Singleton struct {
}

// NewSingleton NewInstance 單例模式方法
func NewSingleton() Singleton {
	once.Do(func() {
		instance = Singleton{}
	})
	return instance
}

var (
	once     sync.Once
	instance Singleton
)

func main() {
	// caller
	s1 := NewSingleton()
	s2 := NewSingleton()
	s3 := NewSingleton()
	fmt.Println(reflect.TypeOf(s1))
	fmt.Println(reflect.TypeOf(s2))
	fmt.Println(reflect.TypeOf(s3))

	// 下面三个打印的地埴都是一样的
	fmt.Printf("%v %p\n", s1, &s1)
	fmt.Printf("%v %p\n", s2, &s2)
	fmt.Printf("%v %p\n", s3, &s3)
}
