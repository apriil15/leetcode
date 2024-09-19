package main

func main() {

}

// time: O(n)
// space: O(n)
func topKFrequent(nums []int, k int) []int {
	numToCount := make(map[int]int)
	for _, num := range nums {
		numToCount[num]++
	}

	// index: count
	// value: nums (because maybe the same count)
	// len should be bigger than len(nums) to prevent out of bound
	countToNums := make([][]int, len(nums)+1)
	for num, count := range numToCount {
		countToNums[count] = append(countToNums[count], num)
	}

	res := make([]int, 0, k)
	for i := len(countToNums) - 1; i >= 0; i-- {
		if len(countToNums[i]) == 0 {
			continue
		}

		res = append(res, countToNums[i]...)
		if len(res) == k {
			return res
		}
	}
	return res
}
