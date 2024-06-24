package main

import "strings"

func main() {

}

// time: O(n^2)
// space: O(1)
func countPrefixSuffixPairs(words []string) int {
	var res int
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if isPrefixAndSuffix(words[i], words[j]) {
				res++
			}
		}
	}

	return res
}

func isPrefixAndSuffix(str1, str2 string) bool {
	return strings.HasPrefix(str2, str1) &&
		strings.HasSuffix(str2, str1)
}
