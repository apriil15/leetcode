package main

import (
	"math"
)

func main() {

}

// time: O(n)
// space: O(1)
func firstMissingPositive(nums []int) int {
	// negative to 0
	for i, num := range nums {
		if num < 0 {
			nums[i] = 0
		}
	}

	// specialNumber can be any number out of (1 ~ len(nums)) * -1
	specialNumber := -(len(nums) + 1)

	// mark negative
	for _, num := range nums {
		index := int(math.Abs(float64(num)) - 1)

		// in bound
		if index >= 0 && index < len(nums) {
			if nums[index] > 0 {
				nums[index] *= -1
			} else if nums[index] == 0 {
				nums[index] = specialNumber
			}
		}
	}

	// check negative
	for i, num := range nums {
		if num >= 0 {
			return i + 1
		}
	}

	// when all negative
	return len(nums) + 1
}
