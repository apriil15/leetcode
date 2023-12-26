package main

import "fmt"

func main() {
	result := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	fmt.Println(result)
}

// time: O(n)
// space: O(1)
func maxSubArray(nums []int) int {
	result := nums[0]
	current := nums[0]

	for i := 1; i < len(nums); i++ {
		if current+nums[i] > nums[i] {
			current = current + nums[i]
		} else {
			current = nums[i]
		}

		if current > result {
			result = current
		}
	}

	return result
}
