package main

import "sort"

func main() {

}

// time: O(n * nlogn)
// space: O(n)
func groupAnagrams_normal(strs []string) [][]string {
	m := make(map[string][]string)

	for _, str := range strs {
		bs := []byte(str)
		sort.Slice(bs, func(i, j int) bool {
			return bs[i] < bs[j]
		})

		if _, ok := m[string(bs)]; ok {
			m[string(bs)] = append(m[string(bs)], str)
		} else {
			m[string(bs)] = []string{str}
		}
	}

	result := make([][]string, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}

	return result
}

// time: O(m*n) (m: len(strs), n: len(str[i]))
func groupAnagrams_tricky(strs []string) [][]string {
	m := make(map[[26]byte][]string)
	for _, str := range strs {
		bs := toKey(str)
		m[bs] = append(m[bs], str)
	}

	result := make([][]string, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}

	return result
}

// eat, tea -> a: 1, e: 1, t: 1
// use [26]byte to represent with (char - 'a')
// example:
// abc -> [1, 1, 1, 0, ..., 0]
func toKey(str string) [26]byte {
	var result [26]byte
	for i := 0; i < len(str); i++ {
		result[str[i]-'a']++
	}

	return result
}
