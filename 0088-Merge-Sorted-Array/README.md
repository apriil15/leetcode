# 88. Merge Sorted Array

## 題目

[Merge Sorted Array - LeetCode](https://leetcode.com/problems/merge-sorted-array/)

## 思路

兩個 slice 從最後一個一直比較，大的就往後丟

比如說：

１　３　７

２　４　５

---

７　跟　５　比較

＿　＿　＿　＿　＿　７

３　跟　５　比較

＿　＿　＿　＿　５　７

３　跟　４　比較

＿　＿　＿　４　５　７

３　跟　２　比較

＿　＿　３　４　５　７

１　跟　２　比較

＿　２　３　４　５　７

最後

１　２　３　４　５　７

## 解法

```go
package main

import (
	"fmt"
)

func main() {
	arr := []int{4, 9, 10, 0, 0, 0}
	arr2 := []int{2, 3, 6}
	merge(arr, 3, arr2, 3)
	fmt.Println(arr)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	// 最大 index
	i := m - 1
	j := n - 1
	k := m + n - 1

	// 概念就是兩個 slice 從最後一個開始比較
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[(i)]
			k--
			i--
		} else {
			nums1[k] = nums2[j]
			k--
			j--
		}
	}

	//代表 nums2 剩下的都比 nums1 剩下的小
	for j >= 0 {
		nums1[k] = nums2[j]
		k--
		j--
	}
}
```
