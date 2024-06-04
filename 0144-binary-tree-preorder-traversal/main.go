package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// mid left right
func preorderTraversal(root *TreeNode) []int {
	var result []int

	var preorder func(root *TreeNode)
	preorder = func(root *TreeNode) {
		if root == nil {
			return
		}

		result = append(result, root.Val)

		preorder(root.Left)
		preorder(root.Right)
	}

	preorder(root)
	return result
}
