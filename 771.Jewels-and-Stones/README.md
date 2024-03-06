# 771. Jewels and Stones

## 題目

[Jewels and Stones - LeetCode](https://leetcode.com/problems/jewels-and-stones/)

## 思路

1. 利用 jewels 做成 map
2. 遍歷 stones 去比對 map

## 解法

```go
package main

import (
	"fmt"
)

func main() {
	jewels := "aA"
	stones := "aAAbbbb"
	result := numJewelsInStones(jewels, stones)
	fmt.Println(result)
}

func numJewelsInStones(jewels string, stones string) int {
	count := 0
	m := make(map[rune]int)
	for _, jewel := range jewels {
		m[jewel] = 0
	}
	for _, stone := range stones {
		if _, ok := m[stone]; ok {
			count++
		}
	}
	return count
}
```