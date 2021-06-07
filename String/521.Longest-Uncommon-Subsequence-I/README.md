# 521. Longest Uncommon Subsequence I

## 題目

[Longest Uncommon Subsequence I - LeetCode](https://leetcode.com/problems/longest-uncommon-subsequence-i/)

大意：找出一個 string 的 subsequence 但不是另一個 string 的 subsequence，而且要找出長度最長的。

## 思路

1. 有種被耍了的感覺啊 XD
2. 如果是`兩個 string 一樣`，找不到 subsequence，回傳 -1
3. 如果`有一個 string 長度較長`，則回傳它的長度

   理由很簡單，較長的 string，它本身（就是自己的 subsequence）一定不是較短 string 的 subsequence

4. 如果`兩個 string 等長，且不是一樣的 string`，直接回傳長度

   理由也很簡單，既然不是一樣的 string，代表兩者至少有一處不同，那麼 a 本身一定不是 b 的 subsequence，反之亦然

## 解法

```go
package main

import "fmt"

func main() {
	a := "aba"
	b := "adc"
	result := findLUSlength(a, b)
	fmt.Println(result)
}

func findLUSlength(a string, b string) int {
	if a == b {
		return -1
	}
	if len(a) > len(b) {
		return len(a)
	}
	return len(b)
}
```
