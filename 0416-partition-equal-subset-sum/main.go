package main

func main() {

}

// time: O(len(nums) * sum count)
// space: O(sum count)
func canPartition(nums []int) bool {
	var sum int
	for _, n := range nums {
		sum += n
	}

	if sum%2 == 1 {
		return false
	}

	target := sum / 2

	// record all possible sum
	set := map[int]struct{}{
		0: {},
	}

	for i := len(nums) - 1; i >= 0; i-- {
		tmpSet := make(map[int]struct{})
		for key := range set {
			t := key + nums[i]
			if t == target {
				return true
			}

			tmpSet[key] = struct{}{}
			tmpSet[t] = struct{}{}
		}
		set = tmpSet
	}
	return false
}
