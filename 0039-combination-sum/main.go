package main

func main() {

}

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var combination []int
	var sum int

	var backtrack func(start int)
	backtrack = func(start int) {
		if sum > target {
			return
		}
		if sum == target {
			dest := make([]int, len(combination))
			copy(dest, combination)
			res = append(res, dest)
			return
		}

		for i := start; i < len(candidates); i++ {
			sum += candidates[i]
			combination = append(combination, candidates[i])

			backtrack(i)

			sum -= candidates[i]
			combination = combination[:len(combination)-1]
		}
	}

	backtrack(0)
	return res
}
