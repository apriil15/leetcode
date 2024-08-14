package main

func main() {

}

// time: O(V+E)
func findOrder(numCourses int, prerequisites [][]int) []int {
	courseToPres := make(map[int][]int) // adjacency list
	for _, pre := range prerequisites {
		courseToPres[pre[0]] = append(courseToPres[pre[0]], pre[1])
	}

	cycle := make(map[int]struct{}) // cycle detect in each dfs
	visit := make(map[int]struct{}) // has been added to res

	var res []int

	var dfs func(course int) bool
	dfs = func(course int) bool {
		if _, ok := cycle[course]; ok {
			return false
		}
		if _, ok := visit[course]; ok {
			return true
		}

		cycle[course] = struct{}{}
		for _, pre := range courseToPres[course] {
			if !dfs(pre) {
				return false
			}
		}
		delete(cycle, course)

		visit[course] = struct{}{}
		res = append(res, course)
		return true
	}

	for i := range numCourses {
		if !dfs(i) {
			return nil
		}
	}
	return res
}
