package main

import "fmt"

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

func subsets(nums []int) [][]int {
	result := [][]int{{}}

	var recursion func(start int, sub []int)
	recursion = func(start int, sub []int) {
		for i := start; i < len(nums); i++ {
			sub = append(sub, nums[i])

			// copy sub to dest
			dest := make([]int, 0, len(sub))
			dest = append(dest, sub...)

			result = append(result, dest)

			recursion(i+1, sub) // go to next layer

			sub = sub[:len(sub)-1]
		}
	}

	recursion(0, []int{})
	return result
}
