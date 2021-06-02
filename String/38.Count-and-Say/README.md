# 38. Count and Say

## 題目

[Count and Say - LeetCode](https://leetcode.com/problems/count-and-say/)

## 思路

1. 要先看懂題目
2. 一層一層推導下去

   `21` → 讀作 `1 個 2，1 個 1` → `1211`

   `1211` → 讀作 `1 個 1，1 個 2，2 個 1` → `111221`

   ```
   1.  1 (base)
   2.  11
   3.  21
   4.  1211
   5.  111221
   6.  312211
   7.  13112221
   8.  1113213211
   9.  31131211131221
   10. 13211311123113112211
   ```

## 解法

```go
package main

import (
	"fmt"
)

func main() {
	result := countAndSay(4)
	fmt.Println(result) // 1211
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	result := "1"
	for i := 1; i < n; i++ {
		result = say(result)
	}
	return result
}

func say(str string) string {
	result := ""
	count := 1

	for i := 1; i < len(str); i++ {
		if str[i-1] == str[i] {
			count++
		} else {
			// 遇到不同數字時，把上一個數字寫進去
			result += fmt.Sprint(count)
			result += string(str[i-1])

			// 重製
			count = 1
		}
	}
	// 處理最後一個數字
	result += fmt.Sprint(count)
	result += string(str[len(str)-1])

	return result
}
```
