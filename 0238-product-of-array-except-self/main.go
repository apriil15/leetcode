package main

func main() {

}

// time: O(n)
// space: O(1)
func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	res[0] = 1

	// calculate left part
	// nums: 2 3  4    5
	// res : 1 2 2*3 2*3*4
	left := 1
	for i := 1; i < len(nums); i++ {
		left *= nums[i-1]
		res[i] = left
	}

	// calculate right part
	// nums:   2    3  4 5
	// res : 3*4*5 4*5 5 1
	right := 1
	for i := len(nums) - 2; i >= 0; i-- {
		right *= nums[i+1]
		res[i] *= right
	}

	return res
}
