package main

func numIslands_bfs(grid [][]byte) int {
	rows := len(grid)
	cols := len(grid[0])
	visit := make(map[point]struct{})
	directions := []point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	var bfs func(i, j int)
	bfs = func(i, j int) {
		q := []point{{i, j}}
		visit[point{i, j}] = struct{}{}

		for len(q) > 0 {
			first := q[0]
			q = q[1:]

			for _, dir := range directions {
				row := first.row + dir.row
				col := first.col + dir.col

				if row < 0 || col < 0 || row >= rows || col >= cols {
					continue
				}
				if _, ok := visit[point{row, col}]; ok {
					continue
				}
				if grid[row][col] == '0' {
					continue
				}

				q = append(q, point{row, col})
				visit[point{row, col}] = struct{}{}
			}
		}
	}

	var res int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if _, ok := visit[point{i, j}]; !ok && grid[i][j] == '1' {
				res++
				bfs(i, j)
			}
		}
	}

	return res
}
