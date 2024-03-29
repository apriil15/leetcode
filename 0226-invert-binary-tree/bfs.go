package main

// BFS (queue)
// time: O(n)
// space: O(n)
func invertTree_bfs(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		first := queue[0]
		queue = queue[1:]

		// swap
		tmp := first.Left
		first.Left = first.Right
		first.Right = tmp

		if first.Left != nil {
			queue = append(queue, first.Left)
		}
		if first.Right != nil {
			queue = append(queue, first.Right)
		}
	}

	return root
}
