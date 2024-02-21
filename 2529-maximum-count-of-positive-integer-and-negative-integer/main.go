package main

import "fmt"

func main() {

	fmt.Println(maximumCount([]int{-2, -1, -1, 1, 2, 3}))
}

// time: O(log n)
// space: O(1)
func maximumCount(nums []int) int {
	// all negative or all positive
	if nums[len(nums)-1] < 0 || nums[0] > 0 {
		return len(nums)
	}

	// all 0
	if nums[0] == 0 && nums[len(nums)-1] == 0 {
		return 0
	}

	// find minPositive
	minPositive := -1
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2

		if mid >= 1 && nums[mid] > 0 && nums[mid-1] <= 0 {
			minPositive = mid
			break
		}

		if nums[mid] <= 0 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	// find maxNegative
	maxNegative := -1
	left = 0
	right = len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2

		if mid <= len(nums)-1 && nums[mid] < 0 && nums[mid+1] >= 0 {
			maxNegative = mid
			break
		}

		if nums[mid] >= 0 {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if minPositive == -1 && maxNegative == -1 {
		return 0
	}

	positiveCount := len(nums) - minPositive
	negativeCount := maxNegative + 1

	if minPositive == -1 {
		return negativeCount
	}
	if maxNegative == -1 {
		return positiveCount
	}

	if positiveCount > negativeCount {
		return positiveCount
	} else {
		return negativeCount
	}
}
