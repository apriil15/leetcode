package main

func main() {

}

// time: O(m*n)
func pacificAtlantic(heights [][]int) [][]int {
	rows := len(heights)
	cols := len(heights[0])

	type point struct {
		row int
		col int
	}
	pac := make(map[point]struct{})
	atl := make(map[point]struct{})

	var dfs func(i, j int, visit map[point]struct{}, preValue int)
	dfs = func(i, j int, visit map[point]struct{}, preValue int) {
		if i < 0 || j < 0 || i >= rows || j >= cols ||
			heights[i][j] < preValue {
			return
		}
		if _, ok := visit[point{i, j}]; ok {
			return
		}

		visit[point{i, j}] = struct{}{}

		dfs(i+1, j, visit, heights[i][j])
		dfs(i-1, j, visit, heights[i][j])
		dfs(i, j+1, visit, heights[i][j])
		dfs(i, j-1, visit, heights[i][j])
	}

	for i := 0; i < rows; i++ {
		dfs(i, 0, pac, 0)      // ->
		dfs(i, cols-1, atl, 0) // <-
	}
	for j := 0; j < cols; j++ {
		dfs(0, j, pac, 0)      // down
		dfs(rows-1, j, atl, 0) // up
	}

	var res [][]int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// find intersection
			if _, ok := pac[point{i, j}]; ok {
				if _, ok := atl[point{i, j}]; ok {
					res = append(res, []int{i, j})
				}
			}
		}
	}
	return res
}
