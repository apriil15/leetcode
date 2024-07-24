package main

func main() {

}

func permuteUnique(nums []int) [][]int {
	var res [][]int
	var per []int

	numToCount := make(map[int]int)
	for _, num := range nums {
		numToCount[num]++
	}

	var dfs func()
	dfs = func() {
		if len(per) == len(nums) {
			dest := make([]int, len(per))
			copy(dest, per)
			res = append(res, dest)
			return
		}

		for num, count := range numToCount {
			if count > 0 {
				per = append(per, num)
				numToCount[num]--

				dfs()

				per = per[:len(per)-1]
				numToCount[num]++
			}
		}
	}

	dfs()
	return res
}
