package main

import "math"

func main() {

}

// time: O(m+n)
// space: O(52) -> O(1)
func minWindow(s string, t string) string {
	// edge case
	if len(t) > len(s) {
		return ""
	}
	if s == t {
		return s
	}

	// tCount to record how much char needed in sliding window
	// + : still need
	// 0 : enough
	// - : too much
	tCount := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tCount[t[i]]++
	}

	// sliding window
	counter := len(t) // counter to check if sliding window is valid
	left, right := 0, 0

	// record for result
	minLength := math.MaxInt
	minLeft := 0

	for right < len(s) {
		// expand right pointer
		if tCount[s[right]] > 0 {
			counter--
		}
		tCount[s[right]]--
		right++

		// when sliding window is valid
		for counter == 0 {
			// calculate result
			if right-left < minLength {
				minLength = right - left
				minLeft = left
			}

			// shrink left pointer
			tCount[s[left]]++
			if tCount[s[left]] > 0 {
				counter++
			}
			left++
		}
	}

	// edge case like:
	// s: "a"
	// t: "b"
	if minLength == math.MaxInt {
		return ""
	}

	return s[minLeft : minLeft+minLength]
}
