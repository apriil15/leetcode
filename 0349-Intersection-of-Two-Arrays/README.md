# 349. Intersection of Two Arrays

## 題目

[Intersection of Two Arrays - LeetCode](https://leetcode.com/problems/intersection-of-two-arrays/)

## 思路

- 解法一

1. 遍歷第一個 slice 建 map
2. 遍歷第二個 slice，如果有在 map 裡面，value++
3. 遍歷 map，如果 value > 0 的，則代表是交集，就把 key 加進去 result

- 解法二（較好）

1. 解法一的步驟三省略，其實在步驟二就可以處理完
2. 步驟二多加一個判斷，只有在第一次交集做處理

## 解法一

```go
package main

import "fmt"

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	result := intersection(nums1, nums2)
	fmt.Println(result)
}

func intersection(nums1 []int, nums2 []int) []int {
	var result []int
	m := make(map[int]int)

	for _, num := range nums1 {
		m[num] = 0
	}

	for _, num := range nums2 {
		if _, ok := m[num]; ok {
			m[num]++
		}
	}

	for key, value := range m {
		if value > 0 {
			result = append(result, key)
		}
	}
	return result
}
```

## 解法二（較好）

```go
package main

import "fmt"

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	result := intersection(nums1, nums2)
	fmt.Println(result)
}

func intersection(nums1 []int, nums2 []int) []int {
	var result []int
	m := make(map[int]int)

	for _, num := range nums1 {
		m[num] = 0
	}

	for _, num := range nums2 {
		if value, ok := m[num]; ok {
			if value == 0 { // 進到這邊代表有交集，但只需要在第一次做處理
				m[num]++
				result = append(result, num)
			}
		}
	}
	return result
}
```
