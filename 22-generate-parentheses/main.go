package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(4))
}

func generateParenthesis(n int) []string {
	result := []string{}

	// open bracket count, close bracket count
	var recursion func(open int, close int, str string)

	recursion = func(open int, close int, str string) {
		if open == n && open == close {
			result = append(result, str)
			return
		}

		if open < n {
			recursion(open+1, close, str+"(")
		}

		if open > close {
			recursion(open, close+1, str+")")
		}
	}

	recursion(0, 0, "")

	return result
}
