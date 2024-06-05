package main

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

	var tmp []int
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		last := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		tmp = append(tmp, last.Val)

		// like preorder, but opposite
		if last.Left != nil {
			stack = append(stack, last.Left)
		}
		if last.Right != nil {
			stack = append(stack, last.Right)
		}
	}

	// reverse
	var res []int
	for i := len(tmp) - 1; i >= 0; i-- {
		res = append(res, tmp[i])
	}

	return res
}
