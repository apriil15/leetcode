package main

import "slices"

func main() {

}

func combinationSum2(candidates []int, target int) [][]int {
	slices.Sort(candidates)

	var res [][]int
	var comb []int
	var sum int

	var backtrack func(start int)
	backtrack = func(start int) {
		if sum == target {
			dest := make([]int, len(comb))
			copy(dest, comb)
			res = append(res, dest)
			return
		}
		if sum > target {
			return
		}

		for i := start; i < len(candidates); i++ {
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}

			comb = append(comb, candidates[i])
			sum += candidates[i]

			backtrack(i + 1)

			comb = comb[:len(comb)-1]
			sum -= candidates[i]
		}
	}

	backtrack(0)
	return res
}
