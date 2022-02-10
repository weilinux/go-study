package main

import "fmt"

func main() {
	var a1 [3]int
	fmt.Println(a1)

	var a2 [3]string
	fmt.Println(a2)

	var a3 [3]float64
	fmt.Println(a3)

	var a4 [3]bool
	fmt.Println(a4)

	for k, v := range a2 {
		fmt.Printf("%v => %v\n", k, v)
	}

	var a5 [3][2]int
	fmt.Println(a5)

	var a6 = [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(a6)
	for key, value := range a6 {
		for k, v := range value {
			fmt.Printf("%v ===> %v => %v\n", key, k, v)
		}
	}

	a7 := [5]int{1, 3, 5, 7, 8}
	var arraySum int
	for _, v := range a7 {
		arraySum += v
	}
	fmt.Println(arraySum)

	for i := 0; i < len(a7); i++ {
		for j := 0; j < len(a7); j++ {
			if a7[i]+a7[j] == 8 {
				fmt.Printf("%v___%v\n", i, j)
			}
		}
	}
}
