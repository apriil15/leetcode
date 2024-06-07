package main

// BFS
// each level declare a "value", update the value when traversing,
// append the "value" after traverse the level
func rightSideView_assign(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		n := len(queue)
		var value int
		for i := 0; i < n; i++ {
			first := queue[0]
			queue = queue[1:]

			value = first.Val

			if first.Left != nil {
				queue = append(queue, first.Left)
			}
			if first.Right != nil {
				queue = append(queue, first.Right)
			}
		}

		res = append(res, value)
	}

	return res
}
