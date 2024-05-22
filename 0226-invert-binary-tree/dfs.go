package main

// DFS (stack)
// time: O(n)
// space: O(n)
func invertTree_dfs(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	stack := []*TreeNode{root}

	for len(stack) > 0 {
		// pop last
		last := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// swap
		tmp := last.Left
		last.Left = last.Right
		last.Right = tmp

		if last.Left != nil {
			stack = append(stack, last.Left)
		}
		if last.Right != nil {
			stack = append(stack, last.Right)
		}
	}

	return root
}
