package main

func main() {

}

// time: O(log n)
// space: O(1)
func findMin(nums []int) int {
	res := nums[0]
	left := 0
	right := len(nums) - 1
	for left <= right {
		if nums[left] <= nums[right] {
			return min(res, nums[left])
		}

		mid := left + (right-left)/2
		res = min(res, nums[mid])

		if nums[left] <= nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return res
}
