package main

func main() {

}

// time: O(26n)
// space: O(26)
func characterReplacement(s string, k int) int {
	m := make(map[byte]int, 26)
	result := 0
	left := 0

	for right := 0; right < len(s); right++ {
		m[s[right]]++

		for (right-left+1)-getMaxCount(m) > k {
			m[s[left]]--
			left++
		}

		if (right - left + 1) > result {
			result = right - left + 1
		}
	}

	return result
}

// time: O(26)
func getMaxCount(m map[byte]int) int {
	result := 0
	for _, n := range m {
		if n > result {
			result = n
		}
	}

	return result
}
