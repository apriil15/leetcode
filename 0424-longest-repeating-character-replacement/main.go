package main

func main() {

}

// time: O(26n)
// space: O(26)
func characterReplacement(s string, k int) int {
	left := 0
	count := make(map[byte]int, 26) // record char count
	res := 0
	for i := 0; i < len(s); i++ {
		count[s[i]]++

		// (i-left+1)-maxCount(count) means the count need to replace
		for (i-left+1)-maxCount(count) > k {
			count[s[left]]--
			left++
		}

		res = max(res, i-left+1)
	}
	return res
}

// time: O(26)
func maxCount(m map[byte]int) int {
	res := 0
	for _, v := range m {
		res = max(res, v)
	}
	return res
}
