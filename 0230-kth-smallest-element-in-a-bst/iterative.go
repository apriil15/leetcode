package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// inorder: left root right
// iterative
// time: O(k)
// space: O(h)
func kthSmallest_iterative(root *TreeNode, k int) int {
	var res int
	stack := []*TreeNode{}
	cur := root

	for len(stack) > 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if k--; k == 0 {
			return cur.Val
		}

		cur = cur.Right
	}

	// would never be here with the constrains
	return res
}
