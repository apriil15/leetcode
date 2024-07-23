package main

import (
	"cmp"
	"slices"
)

func main() {

}

// time: O(nlogn)
// space: O(n)
func frequencySort(nums []int) []int {
	numToCount := make(map[int]int)
	for _, num := range nums {
		numToCount[num]++
	}

	slices.SortFunc(nums, func(a, b int) int {
		if numToCount[a] == numToCount[b] {
			return -cmp.Compare(a, b)
		}

		return cmp.Compare(numToCount[a], numToCount[b])
	})

	return nums
}
