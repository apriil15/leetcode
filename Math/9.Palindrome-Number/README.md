# 9. Palindrome Number

## 題目

[Palindrome Number - LeetCode](https://leetcode.com/problems/palindrome-number/)

## 思路

1. 把每個數字取出，並用 slice 存起來
2. 用兩個 pointer，一前一後比對

## 解法

```go
package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(121))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	// 把每個數字存起來
	s := []int{}
	for x != 0 {
		digit := x % 10 // 取出數字
		s = append(s, digit)
		x /= 10
	}

	// 一前一後比對
	first := 0
	last := len(s) - 1
	for first < last {
		if s[first] != s[last] {
			return false
		}
		first++
		last--
	}
	return true
}
```
