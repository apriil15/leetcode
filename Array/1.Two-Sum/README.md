# 1. Two Sum

## 題目

[Two Sum - LeetCode](https://leetcode.com/problems/two-sum/)

## 解題思路

1. slice 整個遍歷一次就要能找到
2. 將看過的數字存到 map
3. 看新的數字時，用 target 換算成差距的數字，去看看有沒有在 map 裡，有的話就是解答

## 解法

```go
package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 9, 11}, 13))
}

func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}

	m := make(map[int]int)

	for i, num := range nums {
		gap := target - num

		for key, value := range m {
			if gap == value {
				return []int{key, i}
			}
		}
		m[i] = nums[i]
	}
	return nil
}
```
