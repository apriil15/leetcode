package main

import "slices"

func main() {

}

// time: O(n^2)
// space: O(1)
func threeSum(nums []int) [][]int {
	slices.Sort(nums)

	var res [][]int
	for i := 0; i < len(nums)-2; i++ {
		// handle duplicated number
		// [-4, -1, -1, 0, 1, 2]
		// skip 2nd -1
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target := -nums[i]

		// two sum II
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[left] + nums[right]
			if sum > target {
				right--
				continue
			}
			if sum < target {
				left++
				continue
			}

			res = append(res, []int{nums[i], nums[left], nums[right]})

			// handle duplicated number
			// input: [-2, 0, -1, 2, 2]
			// below are the same
			//        [-2, 0,     2]
			//        [-2, 0,        2]
			for left < right {
				right--
				if nums[right] != nums[right+1] {
					break
				}
			}
		}
	}
	return res
}
