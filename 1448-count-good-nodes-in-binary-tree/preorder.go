package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// preorder: root left right
// recursive
// time: O(n)
// space: O(h), O(log n)
func goodNodes_preorder(root *TreeNode) int {
	var res int

	var dfs func(root *TreeNode, maxValue int)
	dfs = func(root *TreeNode, maxValue int) {
		if root == nil {
			return
		}

		if root.Val >= maxValue {
			res++
		}

		dfs(root.Left, max(root.Val, maxValue))
		dfs(root.Right, max(root.Val, maxValue))
	}

	dfs(root, root.Val)
	return res
}
