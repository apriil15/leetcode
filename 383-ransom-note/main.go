package main

func main() {

}

// time: O(m + n)
// space: O(26) -> O(1)
func canConstruct_map(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	// use a map to record magazine
	m := make(map[byte]int, 26)
	for i := 0; i < len(magazine); i++ {
		m[magazine[i]]++
	}

	// if ransomNote's char is in map, then --
	for i := 0; i < len(ransomNote); i++ {
		if _, ok := m[ransomNote[i]]; !ok {
			return false
		} else {
			m[ransomNote[i]]--
			if m[ransomNote[i]] < 0 {
				return false
			}
		}
	}

	return true
}

// time: O(m + n)
// space: O(26) -> O(1)
func canConstruct_slice(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	m := make([]int, 26)
	for i := 0; i < len(magazine); i++ {
		m[magazine[i]-'a']++
	}

	for i := 0; i < len(ransomNote); i++ {
		m[ransomNote[i]-'a']--

		if m[ransomNote[i]-'a'] < 0 {
			return false
		}
	}

	return true
}
