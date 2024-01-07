package main

import (
	"strings"
)

func main() {

}

// time: O(n)
// space: O(n)
func isPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}

	bs := []byte(strings.ToLower(s))

	left := 0
	right := len(bs) - 1
	for left < right {
		if !valid(bs[left]) {
			left++
			continue
		}
		if !valid(bs[right]) {
			right--
			continue
		}

		if bs[left] != bs[right] {
			return false
		}

		left++
		right--
	}

	return true
}

// 0 ~ 9 -> 48 ~ 57
// A ~ Z -> 65 ~ 90
// a ~ z -> 97 ~ 122
func valid(b byte) bool {
	if b >= 48 && b <= 57 {
		return true
	} else if b >= 65 && b <= 90 {
		return true
	} else if b >= 97 && b <= 122 {
		return true
	}

	return false
}
