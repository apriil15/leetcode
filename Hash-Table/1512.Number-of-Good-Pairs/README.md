# 1512. Number of Good Pairs

## 題目

[Number of Good Pairs - LeetCode](https://leetcode.com/problems/number-of-good-pairs/)

## 思路

1. 資料結構用 map

## 解法

- 較好的寫法

```go
package main

import (
	"fmt"
)

func main() {
	result := numIdenticalPairs([]int{1, 2, 3, 1, 1, 3})
	fmt.Println(result)
}

// map 的 key (數字), value (有幾個，從 0 開始計算)
// key:   1 2 3
// value: 2 0 1
func numIdenticalPairs(nums []int) int {
	count := 0
	m := make(map[int]int)

	for _, num := range nums {
		if _, ok := m[num]; ok {
			// 算法的概念：
			// 1 1 1 -> 3 個 pair (1+2)
			// 1 1 1 1 -> 6 個 pair (1+2+3)
			m[num]++
			count += m[num]
		} else {
			m[num] = 0
		}
	}
	return count
}
```

- 比較差的解法

```go
package main

import (
	"fmt"
)

func main() {
	result := numIdenticalPairs([]int{1, 2, 3, 1, 1, 3})
	fmt.Println(result)
}

// map 的 key, value
// key:   0 1 2 3 4 5
// value: 1 2 3 1 1 3
func numIdenticalPairs(nums []int) int {
	count := 0
	m := make(map[int]int)

	for index, num := range nums {
		for _, value := range m {
			if value == num {
				count++
			}
		}
		m[index] = num
	}
	return count
}
```