package main

func main() {

}

// premium problem
// https://neetcode.io/problems/valid-tree
func ValidTree(n int, edges [][]int) bool {
	adj := make(map[int][]int) // undirected adjacency list
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	visit := make(map[int]struct{})

	var dfs func(i, prev int) bool
	dfs = func(i, prev int) bool {
		if _, ok := visit[i]; ok {
			return false
		}

		visit[i] = struct{}{}
		for _, j := range adj[i] {
			if j == prev {
				continue
			}
			if !dfs(j, i) {
				return false
			}
		}

		return true
	}

	return dfs(0, -1) && n == len(visit)
}
