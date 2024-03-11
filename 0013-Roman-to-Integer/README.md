# 13. Roman to Integer

## 題目

[Roman to Integer - LeetCode](https://leetcode.com/problems/roman-to-integer/)

## 思路

1. 先把 map 建好
2. 遍歷 string，把每一個加起來
3. 但如果遇到`前面比後面小時，前面的值要多扣一次`

   舉例：`IV` → `1 + 5 - 1 x 2`

## 解法

```go
package main

import "fmt"

func main() {
	s := "MCMXCIV"
	result := romanToInt(s)
	fmt.Println(result) // 1000 + 100 + 1000 - 100*2 + 10 + 100 - 10*2 + 1 + 5 - 1*2
}

func romanToInt(s string) int {
	m := make(map[rune]int)
	m['I'] = 1
	m['V'] = 5
	m['X'] = 10
	m['L'] = 50
	m['C'] = 100
	m['D'] = 500
	m['M'] = 1000

	result := 0
	tmp := 0
	for _, r := range s {
		if m[r] > tmp { // 如果下一個比較大時，要多扣一次
			result -= tmp * 2
		}
		result += m[r]
		tmp = m[r]
	}
	return result
}
```
