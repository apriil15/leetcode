package main

import "fmt"

func main() {
	result := addDigits(38)
	fmt.Println(result)
}

func addDigits(num int) int {
	for num >= 10 { // 兩位數就要繼續算
		num = addDigit(num)
	}
	return num
}

// 38 -> 11
func addDigit(num int) int {
	result := 0
	for num != 0 {
		digit := num % 10
		result += digit
		num /= 10
	}
	return result
}
