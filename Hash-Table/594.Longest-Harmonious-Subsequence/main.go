package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, 2, 2, 5, 2, 3, 7}
	result := findLHS(nums)
	fmt.Println(result)
}

func findLHS(nums []int) int {
	// 建 map
	// key: 數字, value: 個數
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	result := 0
	for key, value := range m {
		if count, ok := m[key+1]; ok {
			result = max(value+count, result)
		}
	}
	return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
