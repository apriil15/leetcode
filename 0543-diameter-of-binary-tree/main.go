package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// recursion (bottom up)
// time: O(n)
// space: O(n)
func diameterOfBinaryTree(root *TreeNode) int {
	result := 0

	var maxDepth func(root *TreeNode) int
	maxDepth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left := maxDepth(root.Left)
		right := maxDepth(root.Right)

		result = max(result, left+right)

		return max(left, right) + 1
	}

	maxDepth(root)

	return result
}
