package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BFS
// time: O(n)
// space: O(n)
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			first := queue[0]
			queue = queue[1:]

			// only append the last one
			if i == length-1 {
				res = append(res, first.Val)
			}

			if first.Left != nil {
				queue = append(queue, first.Left)
			}
			if first.Right != nil {
				queue = append(queue, first.Right)
			}
		}
	}

	return res
}
