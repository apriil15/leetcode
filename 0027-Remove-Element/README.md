# 27. Remove Element

## 題目

[Remove Element - LeetCode](https://leetcode.com/problems/remove-element/)

## 思路

其實就是遍歷 slice，不是目標數字就把它記下來（原本我們可能直接創一個新的 slice）

但因為題目限制 `不能產生新的 slice`

所以我們直接把值覆蓋上去（題目只要回傳長度內的值對就好）

## 解法

```go
package main

import (
	"fmt"
)

func main() {
	result := removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)
	fmt.Println(result)
}

func removeElement(nums []int, val int) int {
	j := 0
	for _, num := range nums {
		if num != val {
			nums[j] = num
			j++
		}
	}
	return j
}
```