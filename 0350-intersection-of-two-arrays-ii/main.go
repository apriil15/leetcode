package main

func main() {

}

// time: O(n)
// space: O(n)
func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for _, v := range nums1 {
		m[v]++
	}

	var res []int
	for _, num := range nums2 {
		if v, ok := m[num]; ok {
			res = append(res, num)

			v--
			if v == 0 {
				delete(m, num)
			} else {
				m[num]--
			}
		}
	}
	return res
}
