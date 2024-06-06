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
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var res [][]int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		length := len(queue)

		var tmp []int
		for i := 0; i < length; i++ {
			first := queue[0]
			queue = queue[1:]

			tmp = append(tmp, first.Val)

			if first.Left != nil {
				queue = append(queue, first.Left)
			}
			if first.Right != nil {
				queue = append(queue, first.Right)
			}
		}

		res = append(res, tmp)
	}

	return res
}
