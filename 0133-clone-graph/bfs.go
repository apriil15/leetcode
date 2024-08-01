package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */
func cloneGraph_bfs(node *Node) *Node {
	if node == nil {
		return nil
	}

	oldToNew := map[*Node]*Node{
		node: &Node{Val: node.Val},
	}

	queue := []*Node{node}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, neighbor := range cur.Neighbors {
			if _, ok := oldToNew[neighbor]; !ok {
				oldToNew[neighbor] = &Node{Val: neighbor.Val}
				queue = append(queue, neighbor)
			}

			oldToNew[cur].Neighbors = append(
				oldToNew[cur].Neighbors,
				oldToNew[neighbor],
			)
		}
	}

	return oldToNew[node]
}
