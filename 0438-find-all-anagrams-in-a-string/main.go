package main

func main() {

}

// time: O(n)
// space: O(26) -> O(1)
func findAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return []int{}
	}

	pCount := make(map[byte]int, 26)
	for i := 0; i < len(p); i++ {
		pCount[p[i]]++
	}

	result := []int{}

	// let sliding window size the same as input p
	sCount := make(map[byte]int, 26) // char count of sliding window
	for i := 0; i < len(p); i++ {
		sCount[s[i]]++
	}
	if isTheSame(pCount, sCount) {
		result = append(result, 0)
	}

	// start to move sliding window
	left := 0
	for i := len(p); i < len(s); i++ {
		sCount[s[left]]--
		if sCount[s[left]] == 0 {
			delete(sCount, s[left])
		}

		left++
		sCount[s[i]]++

		if isTheSame(pCount, sCount) {
			result = append(result, left)
		}
	}

	return result
}

// time: O(26) -> O(1)
func isTheSame(pCount, sCount map[byte]int) bool {
	for b, count := range pCount {
		if count != sCount[b] {
			return false
		}
	}

	return true
}
