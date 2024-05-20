package main

// BFS (queue)
// imitate maxDepth_dfs_stack()
// time: O(n)
// space: O(n)
func maxDepth_bfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []Node{
		{
			node:  root,
			level: 1,
		},
	}

	result := 1
	for len(queue) > 0 {
		// pop first
		first := queue[0]
		queue = queue[1:]

		result = max(result, first.level)

		if first.node.Left != nil {
			queue = append(queue, Node{
				node:  first.node.Left,
				level: first.level + 1,
			})
		}

		if first.node.Right != nil {
			queue = append(queue, Node{
				node:  first.node.Right,
				level: first.level + 1,
			})
		}
	}

	return result
}

// This one should be better.
// iterate each level, then level++
func maxDepth_bfs2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}

	level := 0
	for len(queue) > 0 {
		// iterate each level
		n := len(queue)
		for i := 0; i < n; i++ {
			// pop first
			first := queue[0]
			queue = queue[1:]

			if first.Left != nil {
				queue = append(queue, first.Left)
			}

			if first.Right != nil {
				queue = append(queue, first.Right)
			}
		}

		level++
	}

	return level
}
