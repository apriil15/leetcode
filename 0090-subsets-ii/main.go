package main

import "slices"

func main() {

}
func subsetsWithDup(nums []int) [][]int {
	slices.Sort(nums) // ASC

	res := [][]int{
		{},
	}
	var subset []int

	var backtrack func(start int)
	backtrack = func(start int) {
		for i := start; i < len(nums); i++ {
			// skip duplicate
			if i > start && nums[i] == nums[i-1] {
				continue
			}

			subset = append(subset, nums[i])

			dest := make([]int, len(subset))
			copy(dest, subset)
			res = append(res, dest)

			backtrack(i + 1)

			subset = subset[:len(subset)-1]
		}
	}

	backtrack(0)
	return res
}
