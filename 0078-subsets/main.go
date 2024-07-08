package main

import "fmt"

func main() {
	fmt.Println(subsets_1([]int{1, 2, 3}))
}

//                     *
//             /       |      \
//            1        2       3
//         /    \      |
//       1,2    1,3   2,3
//       /
//    1,2,3
func subsets_1(nums []int) [][]int {
	res := [][]int{{}}

	var recursion func(start int, sub []int)
	recursion = func(start int, subset []int) {
		for i := start; i < len(nums); i++ {
			subset = append(subset, nums[i])

			// copy subset to dest
			dest := make([]int, 0, len(subset))
			dest = append(dest, subset...)

			res = append(res, dest)

			recursion(i+1, subset) // go to next layer

			subset = subset[:len(subset)-1]
		}
	}

	recursion(0, []int{})
	return res
}

//                    *
//              /           \         add 1 or not
//             1             []
//        /       \       /      \    add 2 or not
//     1,2         1     2       []
//    /   \      /  \   /  \    / \   add 3 or not
// 1,2,3  1,2  1,3  1  2,3  2  3  []
func subsets_2(nums []int) [][]int {
	result := [][]int{}

	var dfs func(i int, sub []int)
	dfs = func(i int, sub []int) {
		if i == len(nums) {
			dest := make([]int, 0, len(sub))
			dest = append(dest, sub...)

			result = append(result, dest)
			return
		}

		sub = append(sub, nums[i])
		dfs(i+1, sub)

		sub = sub[:len(sub)-1]
		dfs(i+1, sub)
	}

	dfs(0, []int{})
	return result
}

// this solution is like the first one,
// but append [] in the backtrack(), and treat subset as a global variable
//
// time: O(n * 2^n)
// space: O(n * 2^n)
func subsets_3(nums []int) [][]int {
	var res [][]int
	var subset []int

	var backtrack func(start int)
	backtrack = func(start int) {
		dest := make([]int, len(subset))
		copy(dest, subset)
		res = append(res, dest)

		for i := start; i < len(nums); i++ {
			subset = append(subset, nums[i])
			backtrack(i + 1)
			subset = subset[:len(subset)-1]
		}
	}

	backtrack(0)
	return res
}
