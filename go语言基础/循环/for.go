package main

import "fmt"

func main() {

	// 九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, j*i)
		}
		if i != 9 {
			fmt.Println()
		}
	}
}
