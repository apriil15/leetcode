package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("dvdf"))
}

// time: O(n)
// space: O(n)
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	set := make(map[byte]struct{})
	left := 0
	result := 0

	for i := 0; i < len(s); i++ {
		for {
			if _, ok := set[s[i]]; ok {
				delete(set, s[left])
				left++
			} else {
				break
			}
		}

		set[s[i]] = struct{}{}

		if len(set) > result {
			result = len(set)
		}
	}

	return result
}
