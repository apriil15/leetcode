package main

func main() {

}

// time: O(V+E)
// space: O(V+E)
func canFinish(numCourses int, prerequisites [][]int) bool {
	courseToPres := make(map[int][]int)
	for i := 0; i < len(prerequisites); i++ {
		row := prerequisites[i][0]
		col := prerequisites[i][1]

		courseToPres[row] = append(courseToPres[row], col)
	}

	cycle := make(map[int]struct{})

	var dfs func(course int) bool
	dfs = func(course int) bool {
		if _, ok := cycle[course]; ok {
			return false
		}
		if len(courseToPres[course]) == 0 {
			return true
		}

		cycle[course] = struct{}{}
		for _, pre := range courseToPres[course] {
			if !dfs(pre) {
				return false
			}
		}
		delete(cycle, course)
		courseToPres[course] = nil

		return true
	}

	// maybe multi graph, so check each
	for i := 0; i < numCourses; i++ {
		if !dfs(i) {
			return false
		}
	}
	return true
}
