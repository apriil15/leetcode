package main

func main() {

}

// time: O(n^2)
// space: O(1)
func countSubstrings(s string) int {
	var res int
	for i := 0; i < len(s); i++ { // <- x ->
		res++

		left := i - 1
		right := i + 1
		for left >= 0 && right <= len(s)-1 {
			if s[left] != s[right] {
				break
			}

			res++

			left--
			right++
		}
	}

	for i := 0; i < len(s); i++ { // <- x x ->
		left := i
		right := i + 1
		for left >= 0 && right <= len(s)-1 {
			if s[left] != s[right] {
				break
			}

			res++

			left--
			right++
		}
	}
	return res
}
