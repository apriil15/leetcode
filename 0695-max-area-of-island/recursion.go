package main

func maxAreaOfIsland_recursion(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || j < 0 || i >= rows || j >= cols {
			return 0
		}
		if grid[i][j] == 0 {
			return 0
		}

		grid[i][j] = 0

		return 1 +
			dfs(i+1, j) +
			dfs(i-1, j) +
			dfs(i, j+1) +
			dfs(i, j-1)
	}

	var res int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				res = max(res, dfs(i, j))
			}
		}
	}

	return res
}
