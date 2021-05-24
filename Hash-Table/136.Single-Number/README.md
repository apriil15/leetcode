# 136. Single Number

## 題目

[Single Number - LeetCode](https://leetcode.com/problems/single-number/)

## 思路

1. 建一個 map
2. 遍歷後如果已經有這個 key 就刪掉
3. 最後會剩下一個 key, value，就是答案

## 解法

```go
package main

import "fmt"

func main() {
	nums := []int{4, 1, 2, 1, 2}
	result := singleNumber(nums)
	fmt.Println(result)
}

func singleNumber(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		if _, ok := m[num]; ok {
			delete(m, num)
		} else {
			m[num] = 0
		}
	}

	var result int
	for key := range m {
		result = key
	}
	return result
}
```
