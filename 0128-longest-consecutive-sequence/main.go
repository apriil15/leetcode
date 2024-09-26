package main

func main() {

}

// time: O(n)
// space: O(n)
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	set := make(map[int]struct{})
	for _, num := range nums {
		set[num] = struct{}{}
	}

	// check if num-1 in set
	// not in set -> num is the first one in its sequence
	var firstOneNums []int
	for num := range set {
		if _, ok := set[num-1]; !ok {
			firstOneNums = append(firstOneNums, num)
		}
	}

	// start to count max length
	var res int
	for _, num := range firstOneNums {
		curMax := 0

		for {
			if _, ok := set[num]; ok {
				curMax++
				num++
			} else {
				break
			}
		}
		res = max(res, curMax)
	}
	return res
}
