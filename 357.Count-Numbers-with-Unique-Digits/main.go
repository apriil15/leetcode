package main

import "fmt"

func main() {
	result := countNumbersWithUniqueDigits(2)
	fmt.Println(result)
}

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 10
	}

	result := 10
	for i := 1; i < n; i++ { // 從 1 開始是因為 n >= 2
		count := 9
		for j := 0; j < i; j++ {
			count *= 9 - j
		}
		result += count
	}
	return result
}
