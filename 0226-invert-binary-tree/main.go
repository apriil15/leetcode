package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// recursion
// time: O(n)
// space: O(n)
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	tmp := root.Left

	root.Left = invertTree(root.Right)
	root.Right = invertTree(tmp)

	return root
}
