# 53. Maximum Subarray

## 題目

[Maximum Subarray - LeetCode](https://leetcode.com/problems/maximum-subarray/)

## 思路

其實是一個領土擴張的概念，我吃掉你有變大的話，就吃掉

沒變大的話，我就變成你就好

(-2) + 1 跟 1 做比較，後者大，所以 current 變成 1

1 + (-3) 跟 -3 做比較，前者大，所以 current 變成 1 + (-3)

這樣一直迭代，同時拿 current 跟 max 做比較，較大的成為新的 max

## 解答

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	result := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	fmt.Println(result)
}

func maxSubArray(nums []int) int {
	max := nums[0]
	current := nums[0]

	for i := 1; i < len(nums); i++ {
		current = int(math.Max(float64(current+nums[i]), float64(nums[i])))
		max = int(math.Max(float64(current), float64(max)))
	}
	return max
}
```