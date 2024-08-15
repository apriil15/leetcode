package main

func main() {

}

// premium problem
// https://neetcode.io/problems/count-connected-components
// time: O(V+E)
func CountComponents(n int, edges [][]int) int {
	adj := make(map[int][]int) // undirected adjacency list
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	visit := make(map[int]struct{})

	var dfs func(i int)
	dfs = func(i int) {
		if _, ok := visit[i]; ok {
			return
		}

		visit[i] = struct{}{}
		for _, j := range adj[i] {
			dfs(j)
		}
	}

	var res int
	for i := range n {
		if _, ok := visit[i]; !ok {
			res++
			dfs(i)
		}
	}
	return res
}
