package main

func addUpper(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		result += i
	}

	return result
}
