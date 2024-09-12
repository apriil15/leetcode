package main

func main() {

}

// time: O(char sum of words)
// space: O(26) -> O(1)
func countConsistentStrings(allowed string, words []string) int {
	m := make(map[rune]struct{})
	for _, c := range allowed {
		m[c] = struct{}{}
	}

	var res int
	for _, word := range words {
		if isConsistent(word, m) {
			res++
		}
	}
	return res
}

// isConsistent check whether each char in the word is in the set
func isConsistent(word string, m map[rune]struct{}) bool {
	for _, c := range word {
		if _, ok := m[c]; !ok {
			return false
		}
	}
	return true
}
