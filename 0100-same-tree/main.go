package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// DFS (stack)
// time: O(n)
// space: O(n)
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if (p == nil && q != nil) ||
		(p != nil && q == nil) {
		return false
	}

	pStack := []*TreeNode{p}
	qStack := []*TreeNode{q}

	result := true
	for len(pStack) > 0 && len(qStack) > 0 {
		pLast := pStack[len(pStack)-1]
		pStack = pStack[:len(pStack)-1]

		qLast := qStack[len(qStack)-1]
		qStack = qStack[:len(qStack)-1]

		if (pLast.Val != qLast.Val) ||
			(pLast.Left != nil && qLast.Left == nil) ||
			(pLast.Left == nil && qLast.Left != nil) ||
			(pLast.Right != nil && qLast.Right == nil) ||
			(pLast.Right == nil && qLast.Right != nil) {
			return false
		}

		if pLast.Left != nil && qLast.Left != nil {
			pStack = append(pStack, pLast.Left)
			qStack = append(qStack, qLast.Left)
		}
		if pLast.Right != nil && qLast.Right != nil {
			pStack = append(pStack, pLast.Right)
			qStack = append(qStack, qLast.Right)
		}
	}

	return result
}
