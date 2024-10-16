package main

func generateParenthesis_backtracking(n int) []string {
	var res []string
	var comb string

	var backtrack func(openCount, closeCount int)
	backtrack = func(openCount, closeCount int) {
		if openCount > n || closeCount > openCount {
			return
		}
		if len(comb) == n*2 {
			res = append(res, comb)
			return
		}

		comb += "("
		backtrack(openCount+1, closeCount)
		comb = comb[:len(comb)-1]

		comb += ")"
		backtrack(openCount, closeCount+1)
		comb = comb[:len(comb)-1]
	}

	backtrack(0, 0)
	return res
}
