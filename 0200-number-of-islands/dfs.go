package main

func numIslands_dfs(grid [][]byte) int {
	rows := len(grid)
	cols := len(grid[0])
	visit := make(map[point]struct{})
	directions := []point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		stack := []point{{i, j}}
		visit[point{i, j}] = struct{}{}

		for len(stack) > 0 {
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			for _, dir := range directions {
				row := last.row + dir.row
				col := last.col + dir.col

				if row < 0 || col < 0 || row >= rows || col >= cols {
					continue
				}
				if _, ok := visit[point{row, col}]; ok {
					continue
				}
				if grid[row][col] == '0' {
					continue
				}

				stack = append(stack, point{row, col})
				visit[point{row, col}] = struct{}{}
			}
		}
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
