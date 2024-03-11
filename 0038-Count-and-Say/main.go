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
