# 3. Longest Substring Without Repeating Characters

## 題目

[Longest Substring Without Repeating Characters - LeetCode](https://leetcode.com/problems/longest-substring-without-repeating-characters/)

## 思路

1. 解法有點像 sliding window (?!)
2. 假設有一個 slice 為 `cdabefa...` 後面還有很多
   遍歷到 `cdabef`，然後下一個看到 `a` 重複了，就把 slice 前面部分的 `cda` 截掉，移動到 `befa`，然後繼續遍歷，找到最長的 slice

## 解法

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	s := "dvdf"
	result := lengthOfLongestSubstring(s)
	fmt.Println(result)
}

func lengthOfLongestSubstring(s string) int {
	max := 0
	arr := []rune{}

	for _, r := range s {
		for i := 0; i < len(arr); i++ {
			if r == arr[i] {
				arr = arr[i+1:] // 遇到重複的 slice 就往右 shift
				break
			}
		}
		arr = append(arr, r)
		max = int(math.Max(float64(max), float64(len(arr))))
	}
	return max
}
```
