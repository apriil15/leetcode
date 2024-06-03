# 594. Longest Harmonious Subsequence

## 題目

[Longest Harmonious Subsequence - LeetCode](https://leetcode.com/problems/longest-harmonious-subsequence/)

## 思路

1. 遍歷 slice 建 map，如下

   數字：`1 3 2 5 7`

   個數：`1 2 3 1 1`

2. 遍歷 map，並找到`下一個數字的個數`（如果存在的話），加總後如果比原本的 result 大，就蓋過去

## 解法

```go
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
```
