package main

func main() {

}

// recursion
func generateParenthesis(n int) []string {
	var res []string

	var recursion func(openCount, closeCount int, comb string)
	recursion = func(openCount, closeCount int, comb string) {
		if openCount > n || closeCount > openCount {
			return
		}
		if len(comb) == n*2 {
			res = append(res, comb)
			return
		}

		recursion(openCount+1, closeCount, comb+"(")
		recursion(openCount, closeCount+1, comb+")")
	}

	recursion(0, 0, "")
	return res
}
