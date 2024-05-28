package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// time: O(m * n)
// space: O(m * n)
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	stack := []*TreeNode{root}

	result := false
	for len(stack) > 0 {
		last := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if isSameTree(last, subRoot) {
			return true
		}

		if last.Left != nil {
			stack = append(stack, last.Left)
		}
		if last.Right != nil {
			stack = append(stack, last.Right)
		}
	}

	return result
}

// time: O(n)
// space: O(n)
func isSameTree(p *TreeNode, q *TreeNode) bool {
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
