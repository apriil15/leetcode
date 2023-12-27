package main

func main() {

}

// https://leetcode.com/problems/contains-duplicate/
// time: O(n)
// space: O(n)
func containsDuplicate(nums []int) bool {
	set := make(map[int]struct{})

	for _, num := range nums {
		if _, ok := set[num]; ok {
			return true
		}

		set[num] = struct{}{}
	}

	return false
}
