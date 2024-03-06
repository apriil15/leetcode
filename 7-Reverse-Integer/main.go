package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(123))
}

func reverse(x int) int {
	result := 0

	// 得到 x 是幾位數
	tmp := x
	length := 0
	for tmp != 0 {
		tmp /= 10
		length++
	}

	count := 1
	for length != 0 {
		digit := x / count % 10 // 3 -> 2 -> 1

		result += digit * int(math.Pow(10, float64(length)-1))
		count *= 10 // 1 -> 10 -> 100
		length--
	}

	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}
	return result
}
