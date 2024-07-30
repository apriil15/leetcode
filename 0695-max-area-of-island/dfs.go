package main

func maxAreaOfIsland_dfs(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	visit := make(map[point]struct{})
	dirs := []point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		stack := []point{{i, j}}
		visit[point{i, j}] = struct{}{}
		count := 1

		for len(stack) > 0 {
			cur := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			for _, dir := range dirs {
				row := cur.row + dir.row
				col := cur.col + dir.col

				if row < 0 || col < 0 || row >= rows || col >= cols {
					continue
				}
				if _, ok := visit[point{row, col}]; ok {
					continue
				}
				if grid[row][col] == 0 {
					continue
				}

				stack = append(stack, point{row, col})
				visit[point{row, col}] = struct{}{}
				count++
			}
		}

		return count
	}

	var res int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if _, ok := visit[point{i, j}]; !ok && grid[i][j] == 1 {
				count := dfs(i, j)
				res = max(res, count)
			}
		}
	}

	return res
}
