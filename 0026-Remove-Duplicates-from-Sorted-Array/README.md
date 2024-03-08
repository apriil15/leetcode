# 26. Remove Duplicates from Sorted Array

## 題目

[Remove Duplicates from Sorted Array - LeetCode](https://leetcode.com/problems/remove-duplicates-from-sorted-array/)

## 解題思路

1. 首先要先看懂題意

   這題要求要在 `同一個 slice`  做操作，最後回傳長度內的元素不能重複，`之後的就不重要了`

- 過程圖

```go
// i 看到不一樣時才會停下來
// j++，然後覆蓋

// j        i
// 1, 1, 1, 2, 4, 5

//    j     i
// 1, 2, 1, 2, 4, 5

//    	 j     i
// 1, 2, 4, 2, 4, 5

//    	    j     i
// 1, 2, 4, 5, 4, 5

// result
// ↓  ↓  ↓  ↓
// 1, 2, 4, 5, 4, 5
```

## 解法

```go
package main

import "fmt"

func main() {
	result := removeDuplicates([]int{1, 1, 1, 2, 4, 5})
	fmt.Println(result)
}

func removeDuplicates(nums []int) int {
	j := 0

	for i := 1; i < len(nums); i++ {
		if nums[j] != nums[i] {
			j++
			nums[j] = nums[i]
		}
	}
	return j + 1
}
```
