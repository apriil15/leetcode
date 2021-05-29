package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"Hello", "Alaska", "Dad", "Peace"}
	result := findWords(words)
	fmt.Println(result)
}

func findWords(words []string) []string {
	// 建立鍵盤的 map
	// 第一列都是 0，第二列都是 1，第三列都是 2
	// ex: q->0, a->1,z->2
	m := make(map[rune]int)
	for index, keyboardRow := range []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"} {
		for _, r := range keyboardRow {
			m[r] = index
		}
	}

	result := []string{}
	for _, word := range words {
		row := 0
		for index, r := range strings.ToLower(word) { // 轉成小寫再比對
			if index == 0 { // 對第一個字母做標記，看它是哪一列
				row = m[r]
				continue
			}
			if row != m[r] { // 不同列就直接 break，換下個字
				break
			}
			if index == len(word)-1 && row == m[r] { // 比對到最後一個且同列，代表這個字都在鍵盤的同一列
				result = append(result, word)
			}
		}
	}
	return result
}
