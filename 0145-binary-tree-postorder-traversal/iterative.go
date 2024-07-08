package main

import "slices"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// left right mid
func postorderTraversal_iterative(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		last := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		res = append(res, last.Val)

		// like preorder, but opposite
		if last.Left != nil {
			stack = append(stack, last.Left)
		}
		if last.Right != nil {
			stack = append(stack, last.Right)
		}
	}

	slices.Reverse(res)

	return res
}
