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
