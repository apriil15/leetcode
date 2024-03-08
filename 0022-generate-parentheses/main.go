package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(generateParenthesis_backtracking(2))
}

func generateParenthesis(n int) []string {
	result := []string{}

	// open bracket count, close bracket count
	var recursion func(open int, close int, str string)

	recursion = func(open int, close int, str string) {
		// base case
		if open == n && close == n {
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

func generateParenthesis_backtracking(n int) []string {
	result := []string{}
	stack := []string{}

	// open bracket count, close bracket count
	var recursion func(open int, close int)

	recursion = func(open int, close int) {
		// base case
		if open == n && close == n {
			result = append(result, strings.Join(stack, ""))
			return
		}

		if open < n {
			stack = append(stack, "(")
			recursion(open+1, close)
			stack = stack[:len(stack)-1] // pop
		}

		if open > close {
			stack = append(stack, ")")
			recursion(open, close+1)
			stack = stack[:len(stack)-1]
		}
	}

	recursion(0, 0)

	return result
}
