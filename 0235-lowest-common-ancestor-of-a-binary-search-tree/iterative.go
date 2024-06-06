package main

// iterative
// time: O(log n)
// space: O(1)
func lowestCommonAncestor_iterative(root, p, q *TreeNode) *TreeNode {
	for {
		if p.Val < root.Val && q.Val < root.Val {
			root = root.Left
		} else if p.Val > root.Val && q.Val > root.Val {
			root = root.Right
		} else {
			return root
		}
	}
}
