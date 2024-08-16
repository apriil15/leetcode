package main

func CountComponents_union_find(n int, edges [][]int) int {
	parents := make([]int, 0, n)
	ranks := make([]int, 0, n)
	for i := range n {
		parents = append(parents, i)
		ranks = append(ranks, 1)
	}

	findParent := func(i int) int {
		res := i

		for res != parents[res] {
			parents[res] = parents[parents[res]] // from neetcode
			res = parents[res]
		}
		return res
	}

	union := func(i, j int) int {
		p1 := findParent(i)
		p2 := findParent(j)

		// already union
		if p1 == p2 {
			return 0
		}

		// union success
		if ranks[p1] > ranks[p2] {
			parents[p2] = p1
			ranks[p1] += ranks[p2]
		} else {
			parents[p1] = p2
			ranks[p2] += ranks[p1]
		}
		return 1
	}

	res := n
	for _, e := range edges {
		res -= union(e[0], e[1])
	}
	return res
}
