package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// iterative
// inorder
// time: O(n)
// space: O(h)
func isValidBST_iterative(root *TreeNode) bool {
	res := true
	stack := []*TreeNode{}
	cur := root
	curMax := math.MinInt

	for len(stack) > 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if cur.Val > curMax {
			curMax = cur.Val
		} else {
			return false
		}

		cur = cur.Right
	}

	return res
}
