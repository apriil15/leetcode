# 728. Self Dividing Numbers

## 題目

[Self Dividing Numbers - LeetCode](https://leetcode.com/problems/self-dividing-numbers/)

## 思路

1. 遍歷每個數字

2. 取出每個位數下去確認

## 解法

```go
package main

import (
	"fmt"
)

func main() {
	result := selfDividingNumbers(1, 22)
	fmt.Println(result)
}

func selfDividingNumbers(left int, right int) []int {
	result := []int{}
	for i := left; i <= right; i++ {
		tmp := i // 複製一份
		for {
			digit := tmp % 10 // 取出最後一位
			if digit == 0 || i%digit != 0 {
				break
			}

			tmp /= 10
			if tmp == 0 {
				result = append(result, i)
				break
			}
		}
	}
	return result
}
```
