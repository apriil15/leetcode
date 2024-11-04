package main

func main() {

}

// time: O(n)
// space: O(n)
func lengthOfLongestSubstring(s string) int {
	left := 0
	res := 0
	set := make(map[byte]struct{})
	for i := 0; i < len(s); i++ {
		// delete until not exist
		for {
			if _, ok := set[s[i]]; ok {
				delete(set, s[left])
				left++
			} else {
				break
			}
		}

		set[s[i]] = struct{}{}
		res = max(res, len(set))
	}
	return res
}
