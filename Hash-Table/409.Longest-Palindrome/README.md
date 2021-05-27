# 409. Longest Palindrome

## 題目

[Longest Palindrome - LeetCode](https://leetcode.com/problems/longest-palindrome/)

## 思路

1. 典型用 map 來解題的題目
2. 數字如果是`偶數`可以直接加進 result
3. `奇數`的話要特別處理，並對 `oddCount++`
   - 假如是 3，代表有 2 個可以回文，所以加上 3-1=2
   - 假如是 1，代表有 0 個可以回文，所以加上 1-1=0
4. 只要 `oddCount` 有一個，就能放在中間，所以 `result++`

## 解法

```go
package main

import (
	"fmt"
)

func main() {
	s := "abccccdd"
	result := longestPalindrome(s)
	fmt.Println(result)
}
func longestPalindrome(s string) int {
	m := make(map[rune]int)
	for _, r := range s {
		m[r]++
	}

	result := 0
	oddCount := 0 // 用來計算有幾個是多出來，不成對的
	for _, value := range m {
		if value%2 == 0 { // 偶數
			result += value
		}
		if value%2 == 1 { // 奇數
			result += value - 1 // 假如是 3，代表有 2 個可以回文；假如是 1，則加上 0
			oddCount++
		}
	}
	if oddCount > 0 { // 只要有一個，就能放在中間，所以 ++
		result++
	}
	return result
}
```
