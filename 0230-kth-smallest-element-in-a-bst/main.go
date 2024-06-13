package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// inorder: left root right
// recursive
// time: O(n)
// space: O(h)
func kthSmallest(root *TreeNode, k int) int {
	var res int

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Left)

		k--
		if k == 0 {
			res = root.Val
		}

		dfs(root.Right)
	}

	dfs(root)
	return res
}
