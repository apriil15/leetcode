package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// time: O(n)
// space: O(n)
func goodNodes(root *TreeNode) int {
	var res int
	stack := []Node{
		{
			TreeNode: root,
			Max:      root.Val,
		},
	}

	for len(stack) > 0 {
		last := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		cur := last.Max

		if last.TreeNode.Val >= last.Max {
			cur = last.TreeNode.Val
			res++
		}

		if last.TreeNode.Left != nil {
			stack = append(stack, Node{
				TreeNode: last.TreeNode.Left,
				Max:      cur,
			})
		}
		if last.TreeNode.Right != nil {
			stack = append(stack, Node{
				TreeNode: last.TreeNode.Right,
				Max:      cur,
			})
		}
	}

	return res
}

type Node struct {
	TreeNode *TreeNode
	Max      int
}
