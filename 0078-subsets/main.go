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
	var subset []int

	var backtrack func(start int)
	backtrack = func(start int) {
		for i := start; i < len(nums); i++ {
			subset = append(subset, nums[i])

			dest := make([]int, len(subset))
			copy(dest, subset)
			res = append(res, dest)

			backtrack(i + 1)

			subset = subset[:len(subset)-1]
		}
	}

	backtrack(0)
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
	var res [][]int
	var subset []int

	var backtrack func(i int)
	backtrack = func(start int) {
		if start == len(nums) {
			dest := make([]int, len(subset))
			copy(dest, subset)
			res = append(res, dest)
			return
		}

		subset = append(subset, nums[start])
		backtrack(start + 1)

		subset = subset[:len(subset)-1]
		backtrack(start + 1)
	}

	backtrack(0)
	return res
}

// this solution is like the first one, but append [] in the backtrack()
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
