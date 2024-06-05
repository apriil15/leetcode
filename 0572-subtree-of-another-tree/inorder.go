package main

// inorder traverse root and check the same tree recursive
func isSubtree_inorder(root *TreeNode, subRoot *TreeNode) bool {
	var result bool

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Left)

		if isSameTree(root, subRoot) {
			result = true
		}

		dfs(root.Right)
	}

	dfs(root)
	return result
}
