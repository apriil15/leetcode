# 35. Search Insert Position

## 題目

[Search Insert Position - LeetCode](https://leetcode.com/problems/search-insert-position/)

給一個排序過的 array，及一個 target，找出 target 可以插入的 index

如果 target 在 array 裡面，就回傳 index

## 思路

1. 題目要求時間複雜度要 `O(log n)`
2. 有排序過
3. 基本上直覺要想到用 `binary search`
4. 有點差異的點為，因為要找插入位置，所以 return left

## 解答

```go
package main

import (
	"fmt"
)

func main() {
	result := searchInsert([]int{1, 3, 5, 6}, 2)
	fmt.Println(result)
}

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid - 1
		}
		if nums[mid] < target {
			left = mid + 1
		}
	}
	return left
}
```