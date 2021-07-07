# 12. Integer to Roman

## 題目

[Integer to Roman - LeetCode](https://leetcode.com/problems/integer-to-roman/)

## 思路

1. 建立一個 struct，裡面有`數字`及`羅馬數字`
2. 建立按順序的對照表
3. 遍歷對照表，機制如下：

   ```
   假設數字是 3994：
   3994 - 1000 = 2994
   2994 - 1000 = 1994
   1994 - 1000 = 994
   994 - 500 = 494
   494 - 400 = 94
   94 - 90 = 4
   4 - 4 = 0
   ```

## 解法

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	result := intToRoman(58)
	fmt.Println(result)
}

func intToRoman(num int) string {
	mapping := []Mapping{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var result strings.Builder // 效能較好
	for _, value := range mapping {
		for num >= value.Integer {
			result.WriteString(value.Roman)
			num -= value.Integer
		}
	}
	return result.String()
}

type Mapping struct {
	Integer int
	Roman   string
}
```
