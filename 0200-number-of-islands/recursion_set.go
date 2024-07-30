package main

func numIslands_recursion_with_set(grid [][]byte) int {
	rows := len(grid)
	cols := len(grid[0])
	visit := make(map[point]struct{})

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || j < 0 || i >= rows || j >= cols {
			return
		}
		if _, ok := visit[point{i, j}]; ok {
			return
		}
		if grid[i][j] == '0' {
			return
		}

		visit[point{i, j}] = struct{}{}

		// up, down, left, right
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}

	var res int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if _, ok := visit[point{i, j}]; !ok && grid[i][j] == '1' {
				res++
				dfs(i, j)
			}
		}
	}

	return res
}

type point struct {
	row int
	col int
}
