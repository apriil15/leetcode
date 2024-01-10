package main

func main() {

}

// time: O(n)
// space: O(26)
func findAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return []int{}
	}

	pMap := make(map[byte]int, 26)
	for i := 0; i < len(p); i++ {
		pMap[p[i]]++
	}

	left := 0
	sMap := make(map[byte]int, 26) // map for sliding window

	result := []int{}

	for i := 0; i < len(s); i++ {
		if count, ok := pMap[s[i]]; !ok {
			left = i + 1
			sMap = make(map[byte]int, 26) // reset map
		} else {
			sMap[s[i]]++

			// shrink left pointer until count valid
			for sMap[s[i]] > count {
				sMap[s[left]]--
				left++
			}
		}

		if i-left+1 == len(p) {
			result = append(result, left)

			sMap[s[left]]--
			left++
		}
	}

	return result
}
