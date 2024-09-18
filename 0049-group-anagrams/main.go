package main

import "slices"

func main() {

}

// time: O(n * nlogn)
// space: O(n)
func groupAnagrams_sort(strs []string) [][]string {
	m := make(map[string][]string)

	for _, str := range strs {
		b := []byte(str)
		slices.Sort(b)
		sortStr := string(b)

		m[sortStr] = append(m[sortStr], str)
	}

	res := make([][]string, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
