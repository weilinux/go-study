package main

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

func main() {
	s := []int{1, 2, 3, 4}
	a := [4]int{5, 6, 7, 8}
	i := 1
	b := []byte{51, 49, 102}
	r := []rune{'阿', '都', '对'}
	s2 := strings.Replace("go go hello", "go", "深圳", 1)
	s3 := strings.TrimSpace(" 阿松 大 ")
	s4 := strings.Trim("s!#das#!fs#s!!s", "#!s")

	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(string(b))
	fmt.Println(string(r))
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(time.Second)
	fmt.Println(time.Millisecond * 100)
}
