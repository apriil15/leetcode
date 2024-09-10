package main

func main() {

}

// time: O(n^2)
// space: O(n)
func lengthOfLIS(nums []int) int {
	// slice to record max subsequence length at each index
	dp := make([]int, len(nums))
	for i := range len(dp) {
		dp[i] = 1
	}

	res := 1
	// the last one always be 1, we traverse from len(nums)-2
	for i := len(nums) - 2; i >= 0; i-- {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
				res = max(res, dp[i])
			}
		}
	}
	return res
}
