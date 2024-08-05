package main

func wallsAndGates_no_visit_set(rooms [][]int) {
	rows := len(rooms)
	cols := len(rooms[0])

	type point struct {
		x int
		y int
	}
	queue := []point{}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if rooms[i][j] == 0 {
				queue = append(queue, point{i, j})
			}
		}
	}

	directions := []point{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, dir := range directions {
			x, y := cur.x+dir.x, cur.y+dir.y
			if x < 0 || y < 0 || x >= rows || y >= cols ||
				rooms[x][y] != 2147483647 {
				continue
			}

			rooms[x][y] = rooms[cur.x][cur.y] + 1
			queue = append(queue, point{x, y})
		}
	}
}
