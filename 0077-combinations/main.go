package main

import "fmt"

func main() {
	fmt.Println(combine(3, 2))
}

func combine(n int, k int) [][]int {
	result := [][]int{}

	var backtrack func(start int, comb []int)
	backtrack = func(start int, comb []int) {
		// base case
		if len(comb) == k {
			// copy comb to dest
			dest := make([]int, 0, len(comb))
			dest = append(dest, comb...)

			result = append(result, dest)
			return
		}

		for i := start; i <= n; i++ {
			comb = append(comb, i)
			backtrack(i+1, comb)      // go to next layer
			comb = comb[:len(comb)-1] // back
		}
	}

	backtrack(1, []int{})
	return result
}
