package main

import "sort"

func main() {

}

// time: O(s+t)
// space: O(1)
func isAnagram_1map(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// space O(1), time O(s)
	charToCount := make(map[rune]int, 26)
	for _, char := range s {
		charToCount[char]++
	}

	// time O(t)
	for _, char := range t {
		if v, ok := charToCount[char]; ok {
			if v > 0 {
				charToCount[char]--
			} else {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

// time: O(nlogn)
// space: O(1)
func isAnagram_sort(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	ss := []rune(s)
	tt := []rune(t)
	sort.Slice(ss, func(i, j int) bool {
		return ss[i] < ss[j]
	})
	sort.Slice(tt, func(i, j int) bool {
		return tt[i] < tt[j]
	})

	for i := 0; i < len(ss); i++ {
		if ss[i] != tt[i] {
			return false
		}
	}

	return true
}

// time: O(len(s)+len(t))
// space: O(52) -> O(1)
func isAnagram_2map(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// map of char count
	m1 := make(map[rune]int, 26)
	for _, char := range s {
		m1[char]++
	}

	m2 := make(map[rune]int, 26)
	for _, char := range t {
		m2[char]++
	}

	for char, count := range m1 {
		c, ok := m2[char]
		if !ok {
			return false
		}
		if count != c {
			return false
		}
	}
	return true
}
