package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// left right mid
func postorderTraversal(root *TreeNode) []int {
	var res []int

	var postorder func(root *TreeNode)
	postorder = func(root *TreeNode) {
		if root == nil {
			return
		}

		postorder(root.Left)
		postorder(root.Right)

		res = append(res, root.Val)
	}

	postorder(root)
	return res
}
