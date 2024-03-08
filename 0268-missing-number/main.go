package main

import (
	"sort"
)

func main() {

}

// time: O(nlogn)
// space: O(1)
func missingNumber_sort(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	result := 0
	for _, num := range nums {
		if num != result {
			return result
		}

		result++
	}

	return result
}

// time: O(n)
// space: O(n)
func missingNumber_set(nums []int) int {
	set := make(map[int]struct{}, len(nums)+1)
	for i := 0; i < len(nums)+1; i++ {
		set[i] = struct{}{}
	}

	for _, num := range nums {
		delete(set, num)
	}

	for k := range set {
		return k
	}

	return -1
}

// time: O(n)
// space: O(1)
func missingNumber_sum(nums []int) int {
	total := (len(nums) + 1) * (0 + len(nums)) / 2

	sum := 0
	for _, num := range nums {
		sum += num
	}

	return total - sum
}
