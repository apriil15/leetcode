# 20. Valid Parentheses

## 題目

[Valid Parentheses - LeetCode](https://leetcode.com/problems/valid-parentheses/)

## 思路

1. 先把配對的 map 建好
2. `左括號`直接加進去
3. `右括號`則判斷能不能跟前一個配對

   可以配對 → 把`前一個刪掉`

   無法配對 → 回傳 false

4. Golang 沒有 stack，因此用 slice 處理

## 解法

```go
package main

import "fmt"

func main() {
	s := "()[]{}"
	result := isValid(s)
	fmt.Println(result)
}

func isValid(s string) bool {
	m := make(map[rune]rune)
	m['('] = ')'
	m['['] = ']'
	m['{'] = '}'

	stack := []rune{}
	for _, r := range s {
		if _, ok := m[r]; ok { // 左邊直接加進去
			stack = append(stack, r)
		} else if len(stack) == 0 { // 開頭為右邊 ")"
			return false
		} else if r != m[stack[len(stack)-1]] { // 經過 map 轉換不同，代表跟上一個無法配對
			return false
		} else {
			stack = stack[:len(stack)-1] // 可以配對，則把最後一個刪掉
		}
	}
	return len(stack) == 0
}
```
