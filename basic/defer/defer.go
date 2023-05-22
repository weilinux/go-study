package main

import "fmt"

func f1() int {
	i := 5
	defer func() {
		i++
	}()
	return i
}

func f2() (i int) {
	i = 5
	defer func() {
		i++
	}()
	return i
}

func main() {
	x := f1()
	fmt.Println(x)

	y := f2()
	fmt.Println(y)
}
