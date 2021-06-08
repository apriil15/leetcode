# 557. Reverse Words in a String III

## 題目

[Reverse Words in a String III - LeetCode](https://leetcode.com/problems/reverse-words-in-a-string-iii/)

## 思路

1. 很直覺地把空格切掉，再對每個 string 做反轉
2. 本來是用 string 直接加，改用 `strings.Builder` 效能會好很多

## 解法

- 效能較佳

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Let's take LeetCode contest"
	result := reverseWords(s)
	fmt.Println(result)
}

func reverseWords(s string) string {
	// 先把空白切掉
	arr := strings.Split(s, " ")

	// 改用 strings.Builder
	var builder strings.Builder

	// 把每個做反轉
	for index, str := range arr {
		for i := len(str) - 1; i >= 0; i-- {
			builder.WriteString(string(str[i]))
		}
		if index == len(arr)-1 { // 最後一個不用加空格
			break
		}
		builder.WriteString(" ")
	}
	return builder.String()
}
```

- 效能很差

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Let's take LeetCode contest"
	result := reverseWords(s)
	fmt.Println(result)
}

func reverseWords(s string) string {
	// 先把空白切掉
	arr := strings.Split(s, " ")

	// 把每個做反轉
	result := ""
	for index, str := range arr {
		for i := len(str) - 1; i >= 0; i-- {
			result += string(str[i])
		}
		if index == len(arr)-1 { // 最後一個不用加空格
			break
		}
		result += " "
	}
	return result
}
```
