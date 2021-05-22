# 119. Pascal's Triangle II

## 題目

[Pascal's Triangle II - LeetCode](https://leetcode.com/problems/pascals-triangle-ii/)

## 思路

用上一層來推導下一層

## 解法

```go
package main

import (
	"fmt"
)

func main() {
	numRows := 3
	result := getRow(numRows)
	fmt.Println(result)
}

func getRow(rowIndex int) []int {
	result := []int{}

	for i := 0; i <= rowIndex; i++ {
		tmp := make([]int, len(result)) // tmp 就是上一層
		copy(tmp, result)
		result = append(result, 1) // 最後一個都是 1

		// 因為頭尾都是 1，所以不用更新，要更新第二個~倒數第二個
		for j := 1; j <= len(result)-2; j++ {
			result[j] = tmp[j-1] + tmp[j]
		}
	}
	return result
}
```