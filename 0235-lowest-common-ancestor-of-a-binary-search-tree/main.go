package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// time: O(log n)
// space: O(depth of n)
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if (root.Val == p.Val || root.Val == q.Val) ||
		(p.Val < root.Val && q.Val > root.Val) ||
		(q.Val < root.Val && p.Val > root.Val) {
		return root
	}

	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}

	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}

	return nil
}
