package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// iterative
// time: O(n)
// space: O(h), O(log n)
func goodNodes(root *TreeNode) int {
	var res int
	stack := []Node{
		{
			node:     root,
			maxValue: root.Val,
		},
	}

	for len(stack) > 0 {
		last := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if last.node.Val >= last.maxValue {
			res++
		}

		maxValue := max(last.node.Val, last.maxValue)

		if last.node.Left != nil {
			stack = append(stack, Node{
				node:     last.node.Left,
				maxValue: maxValue,
			})
		}
		if last.node.Right != nil {
			stack = append(stack, Node{
				node:     last.node.Right,
				maxValue: maxValue,
			})
		}
	}

	return res
}

type Node struct {
	node     *TreeNode
	maxValue int
}
