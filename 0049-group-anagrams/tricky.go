package main

// time: O(m*n) (m: len(strs), n: len(str[i]))
func groupAnagrams_tricky(strs []string) [][]string {
	m := make(map[[26]byte][]string)

	for _, str := range strs {
		charCount := [26]byte{}
		for i := 0; i < len(str); i++ {
			charCount[str[i]-'a']++
		}
		m[charCount] = append(m[charCount], str)
	}

	res := make([][]string, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
