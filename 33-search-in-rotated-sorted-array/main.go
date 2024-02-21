package main

import "fmt"

func main() {
	fmt.Println(search([]int{3, 1}, 1))
}

// time: O(log n)
// space: O(1)
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		// left side is ordered
		if nums[left] <= nums[mid] {
			// target is in this side
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// right side is ordered
			// and target is in this side
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}
