package main

func main() {

}

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	parents := make([]int, 0, n)
	ranks := make([]int, 0, n)
	for i := range n {
		parents = append(parents, i)
		ranks = append(ranks, 1)
	}

	find := func(i int) int {
		res := i
		for res != parents[res] {
			parents[res] = parents[parents[res]]
			res = parents[res]
		}
		return res
	}

	union := func(i, j int) bool {
		p1 := find(i)
		p2 := find(j)

		// same parent -> the same union set
		if p1 == p2 {
			return false
		}

		if ranks[p1] > ranks[p2] {
			parents[p2] = p1
			ranks[p1] += ranks[p2]
		} else {
			parents[p1] = p2
			ranks[p2] += ranks[p1]
		}

		return true
	}

	for _, e := range edges {
		// offset -1 because vertices value starts with 1
		if !union(e[0]-1, e[1]-1) {
			return e
		}
	}

	return nil
}

func findRedundantConnection_without_offset(edges [][]int) []int {
	n := len(edges)
	parents := []int{-1}
	ranks := []int{-1}
	for i := 1; i <= n; i++ {
		parents = append(parents, i)
		ranks = append(ranks, 1)
	}

	find := func(i int) int {
		res := i
		for res != parents[res] {
			parents[res] = parents[parents[res]]
			res = parents[res]
		}
		return res
	}

	union := func(i, j int) bool {
		p1 := find(i)
		p2 := find(j)

		// same parent -> the same union set
		if p1 == p2 {
			return false
		}

		if ranks[p1] > ranks[p2] {
			parents[p2] = p1
			ranks[p1] += ranks[p2]
		} else {
			parents[p1] = p2
			ranks[p2] += ranks[p1]
		}

		return true
	}

	for _, e := range edges {
		if !union(e[0], e[1]) {
			return e
		}
	}

	return nil
}
