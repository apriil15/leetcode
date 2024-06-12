package main

import "math"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// inorder: left root right
// should be larger and larger
// time: O(n)
// space: O(h), O(log n)
func isValidBST(root *TreeNode) bool {
	res := true
	cur := math.MinInt

	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}

		inorder(root.Left)

		if root.Val > cur {
			cur = root.Val
		} else {
			res = false
		}

		inorder(root.Right)
	}

	inorder(root)
	return res
}
