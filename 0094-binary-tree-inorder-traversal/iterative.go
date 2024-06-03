package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal_iterative(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := []int{}
	stack := []*TreeNode{}
	cur := root

	for len(stack) > 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, cur.Val)

		cur = cur.Right
	}

	return result
}
