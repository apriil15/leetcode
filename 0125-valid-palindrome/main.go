package main

import (
	"strings"
)

func main() {

}

// time: O(n)
// space: O(1)
func isPalindrome(s string) bool {
	s = strings.ToLower(s)

	left := 0
	right := len(s) - 1
	for left < right {
		if !isValid(s[left]) {
			left++
			continue
		}
		if !isValid(s[right]) {
			right--
			continue
		}

		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// '0': 48
// 'A': 65
// 'a': 97
func isValid(b byte) bool {
	return (b >= 48 && b <= 57) ||
		(b >= 97 && b <= 122)
}
