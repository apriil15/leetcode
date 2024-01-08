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
		if _, ok := set[s[i]]; !ok {
			set[s[i]] = struct{}{}
		} else {
			for {
				if s[left] == s[i] {
					left++
					break
				}

				delete(set, s[left])
				left++
			}
		}

		if len(set) > result {
			result = len(set)
		}
	}

	return result
}
