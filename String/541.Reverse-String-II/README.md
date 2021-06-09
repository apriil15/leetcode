# 541. Reverse String II

## 題目

[Reverse String II - LeetCode](https://leetcode.com/problems/reverse-string-ii/)

## 思路

1. 想法很直覺，每 `2k 個`裡面的`前 k 個`做 reverse
2. 要注意的是 Golang 的 `swap` 可以如解法這樣實作，相當方便

## 解法

```go
package main

import "fmt"

func main() {
	s := "abcdefg"
	k := 2
	result := reverseStr(s, k)
	fmt.Println(result)
}

func reverseStr(s string, k int) string {
	arr := []byte(s)
	length := len(arr)

	// 因為每 2k 個，只有其中的前面 k 個要 reverse，因此用 bool 當作開關
	shouldReverse := true

	for i := 0; i < length; i += k {
		if shouldReverse {
			end := i + k - 1
			if end > length-1 { // 如果超出最後一個，就改成最後一個
				end = length - 1
			}
			reverse(arr, i, end)
		}
		shouldReverse = !shouldReverse
	}
	return string(arr)
}

func reverse(arr []byte, start int, end int) {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start] // swap
		start++
		end--
	}
}
```
