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

		leftMax := maxDepth(root.Left)
		rightMax := maxDepth(root.Right)

		if math.Abs(float64(leftMax-rightMax)) > 1 {
			result = false
		}

		return max(leftMax, rightMax) + 1
	}

	maxDepth(root)

	return result
}
