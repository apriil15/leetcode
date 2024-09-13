package main

func main() {

}

// time: O(n)
// space: O(n)
func containsDuplicate(nums []int) bool {
	seen := make(map[int]struct{})

	for _, n := range nums {
		if _, ok := seen[n]; ok {
			return true
		}
		seen[n] = struct{}{}
	}
	return false
}
