package main

// Node for maxDepth_dfs_stack
type Node struct {
	node  *TreeNode
	level int
}

// DFS (stack)
// time: O(n)
// space: O(n)
func maxDepth_dfs_stack(root *TreeNode) int {
	if root == nil {
		return 0
	}

	stack := []Node{
		{
			node:  root,
			level: 1,
		},
	}

	result := 1
	for len(stack) > 0 {
		// pop
		last := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = max(result, last.level)

		if last.node.Left != nil {
			stack = append(stack, Node{
				node:  last.node.Left,
				level: last.level + 1,
			})
		}

		if last.node.Right != nil {
			stack = append(stack, Node{
				node:  last.node.Right,
				level: last.level + 1,
			})
		}
	}

	return result
}
