package main

func main() {

}

// time: O(n)
// space: O(1)
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	// skip first vs skip last
	return max(helper(nums[1:]), helper(nums[:len(nums)-1]))
}

// 198. House Robber
func helper(nums []int) int {
	rob1 := 0
	rob2 := 0
	for _, n := range nums {
		tmp := max(n+rob1, rob2)
		rob1 = rob2
		rob2 = tmp
	}
	return rob2
}
