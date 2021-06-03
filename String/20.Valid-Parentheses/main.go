package main

import "fmt"

func main() {
	s := "()[]{}"
	result := isValid(s)
	fmt.Println(result)
}

func isValid(s string) bool {
	m := make(map[rune]rune)
	m['('] = ')'
	m['['] = ']'
	m['{'] = '}'

	stack := []rune{}
	for _, r := range s {
		if _, ok := m[r]; ok { // 左邊直接加進去
			stack = append(stack, r)
		} else if len(stack) == 0 { // 開頭為右邊 ")"
			return false
		} else if r != m[stack[len(stack)-1]] { // 經過 map 轉換不同，代表跟上一個無法配對
			return false
		} else {
			stack = stack[:len(stack)-1] // 可以配對，則把最後一個刪掉
		}
	}
	return len(stack) == 0
}
