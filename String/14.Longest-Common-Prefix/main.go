package main

import "fmt"

func main() {
	strs := []string{"ab", "a"}
	result := longestCommonPrefix(strs)
	fmt.Println(result)
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	// 找到最短的 string 長度
	// ex: {"ab", "a"} 只要看 1 個 char 就行
	minLength := len(strs[0])
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < minLength {
			minLength = len(strs[i])
		}
	}

	result := ""
	for i := 0; i < minLength; i++ {
		char := strs[0][i]               // 把每個 char 丟進來
		for j := 1; j < len(strs); j++ { // 對每個 string 做確認
			if char != strs[j][i] {
				return result
			}
			if char == strs[j][i] && j == len(strs)-1 { // 一樣且是最後一個 string，代表全部都一樣
				result += string(char)
			}
			if char == strs[j][i] { // 這個判斷其實可以不用寫
				continue
			}
		}
	}
	return result
}
