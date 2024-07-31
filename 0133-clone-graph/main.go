package main

func main() {

}

type Node struct {
	Val       int
	Neighbors []*Node
}

// time: O(V+E)
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	oldToNew := make(map[*Node]*Node)

	var dfs func(node *Node) *Node
	dfs = func(node *Node) *Node {
		if v, ok := oldToNew[node]; ok {
			return v
		}

		copy := &Node{Val: node.Val}
		oldToNew[node] = copy

		for _, neighbor := range node.Neighbors {
			copy.Neighbors = append(copy.Neighbors, dfs(neighbor))
		}
		return copy
	}

	return dfs(node)
}
