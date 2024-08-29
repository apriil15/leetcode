package main

func main() {
}

// two pointer
// time: O(n^2)
// space: O(1)
func longestPalindrome(s string) string {
	res := s[:1]

	// use i as center, move left & right pointer
	for i := 0; i < len(s); i++ {
		left := i - 1
		right := i + 1

		for left >= 0 && right <= len(s)-1 {
			if s[left] != s[right] {
				break
			}

			if right-left+1 > len(res) {
				res = s[left : right+1]
			}

			left--
			right++
		}
	}

	// use adjacent as start point, move left & right pointer
	for i := 0; i < len(s); i++ {
		left := i
		right := i + 1

		for left >= 0 && right <= len(s)-1 {
			if s[left] != s[right] {
				break
			}

			if right-left+1 > len(res) {
				res = s[left : right+1]
			}

			left--
			right++
		}
	}

	return res
}
