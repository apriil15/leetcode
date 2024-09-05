package main

func main() {

}

// time: O(n)
// space: O(1)
func maxProduct(nums []int) int {
	res := nums[0]
	curMax := nums[0]
	curMin := nums[0]

	for i := 1; i < len(nums); i++ {
		// swap here because max*negative become min, and min*negative become max
		if nums[i] < 0 {
			curMax, curMin = curMin, curMax
		}

		curMax = max(nums[i], curMax*nums[i])
		curMin = min(nums[i], curMin*nums[i])

		res = max(res, curMax)
	}
	return res
}
