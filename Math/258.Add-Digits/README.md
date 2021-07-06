# 258. Add Digits

## 題目

[Add Digits - LeetCode](https://leetcode.com/problems/add-digits/)

## 思路

1. 把每個數字取出，加起來
2. recursion 持續做，直到是個位數

## 解法

```go
package main

import "fmt"

func main() {
	result := addDigits(38)
	fmt.Println(result)
}

func addDigits(num int) int {
	for num >= 10 { // 兩位數就要繼續算
		num = addDigit(num)
	}
	return num
}

// 38 -> 11
func addDigit(num int) int {
	result := 0
	for num != 0 {
		digit := num % 10
		result += digit
		num /= 10
	}
	return result
}
```
