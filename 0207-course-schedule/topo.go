package main

func canFinish_another_set(numCourses int, prerequisites [][]int) bool {
	courseToPres := make(map[int][]int)
	for _, pre := range prerequisites {
		courseToPres[pre[0]] = append(courseToPres[pre[0]], pre[1])
	}

	cycle := make(map[int]struct{})
	visit := make(map[int]struct{})

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

		return true
	}

	for i := range numCourses {
		if !dfs(i) {
			return false
		}
	}
	return true
}
