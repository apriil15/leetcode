package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// mid left right
func preorderTraversal_iterative(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		last := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		res = append(res, last.Val)

		// push "right" first, then "left"
		if last.Right != nil {
			stack = append(stack, last.Right)
		}
		if last.Left != nil {
			stack = append(stack, last.Left)
		}
	}

	return res
}
