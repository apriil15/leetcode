package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// time: O(n)
// space: O(n)
func getMinimumDifference(root *TreeNode) int {
	result := math.MaxInt
	var previous *int

	var inOrder func(root *TreeNode)
	inOrder = func(root *TreeNode) {
		if root == nil {
			return
		}

		inOrder(root.Left)

		if previous != nil {
			diff := root.Val - *previous
			result = min(result, diff)
		}

		previous = &root.Val

		inOrder(root.Right)
	}

	inOrder(root)
	return result
}
