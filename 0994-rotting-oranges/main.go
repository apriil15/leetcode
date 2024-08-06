package main

func main() {

}

func orangesRotting(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	type point struct {
		row int
		col int
	}
	q := []point{}
	freshCount := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 2 {
				q = append(q, point{i, j})
			} else if grid[i][j] == 1 {
				freshCount++
			}
		}
	}

	if freshCount == 0 {
		return 0
	}

	res := -1
	dirs := []point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for len(q) > 0 {
		length := len(q)
		for i := 0; i < length; i++ {
			cur := q[0]
			q = q[1:]

			for _, dir := range dirs {
				row := cur.row + dir.row
				col := cur.col + dir.col
				if row < 0 || col < 0 || row >= rows || col >= cols {
					continue
				}
				if grid[row][col] != 1 {
					continue
				}

				grid[row][col] = 2
				q = append(q, point{row, col})
				freshCount--
			}
		}
		res++
	}

	if freshCount == 0 {
		return res
	}
	return -1
}
