package main

func wallsAndGates_no_visit_set(rooms [][]int) {
	rows := len(rooms)
	cols := len(rooms[0])

	type point struct {
		row int
		col int
	}
	q := []point{}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if rooms[i][j] == 0 {
				q = append(q, point{i, j})
			}
		}
	}

	dirs := []point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		for _, dir := range dirs {
			row := cur.row + dir.row
			col := cur.col + dir.col
			if row < 0 || col < 0 || row >= rows || col >= cols ||
				rooms[row][col] != 2147483647 {
				continue
			}

			rooms[row][col] = rooms[cur.row][cur.col] + 1
			q = append(q, point{row, col})
		}
	}
}
