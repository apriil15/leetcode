package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	result := math.MaxInt
	var previous *int

	// in-order
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Left)

		if previous != nil {
			diff := root.Val - *previous
			result = min(result, diff)
		}

		previous = &root.Val

		dfs(root.Right)
	}

	dfs(root)
	return result
}
