package main

import "math"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	result := true

	var maxDepth func(root *TreeNode) int
	maxDepth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left := maxDepth(root.Left)
		right := maxDepth(root.Right)

		if math.Abs(float64(left-right)) > 1 {
			result = false
		}

		return max(left, right) + 1
	}

	maxDepth(root)

	return result
}
