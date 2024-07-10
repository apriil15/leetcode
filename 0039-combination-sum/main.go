package main

func main() {

}

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var comb []int
	var sum int

	var backtrack func(start int)
	backtrack = func(start int) {
		if sum > target {
			return
		}
		if sum == target {
			dest := make([]int, len(comb))
			copy(dest, comb)
			res = append(res, dest)
			return
		}

		for i := start; i < len(candidates); i++ {
			comb = append(comb, candidates[i])
			sum += candidates[i]

			backtrack(i)

			comb = comb[:len(comb)-1]
			sum -= candidates[i]
		}
	}

	backtrack(0)
	return res
}
