package main

// recursion
// time: O(n)
// space: O(n)
func isSameTree_recursion(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if (p == nil && q != nil) ||
		(p != nil && q == nil) ||
		(p.Val != q.Val) {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
