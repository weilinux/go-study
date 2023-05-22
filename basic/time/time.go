package main

import (
	"fmt"
	"time"
)

func errorMsg() {
	if p := recover(); p != nil {
		fmt.Println(p)
	}
}

func main() {
	defer errorMsg()

	dateTime := time.Now()
	unix := dateTime.Unix()
	fmt.Println(dateTime)
	fmt.Println(unix)
	fmt.Println(dateTime.Format("2006-01-02 15:04:05"))

	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:4]
	fmt.Println(&a[0])
	fmt.Println(&a[1])
	fmt.Println(&a[2])
	fmt.Println(&s[0])
	fmt.Println(&s[1])

	s2 := make([]int, 5, 10)
	fmt.Println(s2)
	s2 = append(s2, 10, 20, 30)
	fmt.Println(s2)
	fmt.Println(&s2[0])
	copy(s2, s)
	fmt.Println(s2)
	fmt.Println(&s2[0])
	fmt.Println(&s2[0])
	s2[3] = a[3]
	fmt.Println(s2, a)
	fmt.Println(&s2[3], &a[3])
	fmt.Println(&s2[0])

	panic("sadasdasd")
	fmt.Println(1)
}
