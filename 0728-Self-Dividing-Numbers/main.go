package main

import (
	"fmt"
)

func main() {
	result := selfDividingNumbers(1, 22)
	fmt.Println(result)
}

func selfDividingNumbers(left int, right int) []int {
	result := []int{}
	for i := left; i <= right; i++ {
		tmp := i // 複製一份
		for {
			digit := tmp % 10 // 取出最後一位
			if digit == 0 || i%digit != 0 {
				break
			}

			tmp /= 10
			if tmp == 0 {
				result = append(result, i)
				break
			}
		}
	}
	return result
}
