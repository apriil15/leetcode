package main

import "fmt"

func main() {
	rooms := [][]int{
		{2147483647, -1, 0, 2147483647},
		{2147483647, 2147483647, 2147483647, -1},
		{2147483647, -1, 2147483647, -1},
		{0, -1, 2147483647, 2147483647},
	}

	wallsAndGates(rooms)

	for _, row := range rooms {
		fmt.Println(row)
	}
	// [3 -1 0  1]
	// [2  2 1 -1]
	// [1 -1 2 -1]
	// [0 -1 3  4]
}

// premium problem
// https://neetcode.io/problems/islands-and-treasure
func wallsAndGates(rooms [][]int) {
	rows := len(rooms)
	cols := len(rooms[0])

	type point struct {
		row int
		col int
	}
	visit := make(map[point]struct{})
	q := []point{}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if rooms[i][j] == 0 {
				q = append(q, point{i, j})
				visit[point{i, j}] = struct{}{}
			}
		}
	}

	dist := 0
	dirs := []point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for len(q) > 0 {
		length := len(q)
		for i := 0; i < length; i++ {
			cur := q[0]
			q = q[1:]

			rooms[cur.row][cur.col] = dist

			for _, dir := range dirs {
				row := cur.row + dir.row
				col := cur.col + dir.col
				if row < 0 || col < 0 || row >= rows || col >= cols ||
					rooms[row][col] == -1 {
					continue
				}
				if _, ok := visit[point{row, col}]; ok {
					continue
				}

				q = append(q, point{row, col})
				visit[point{row, col}] = struct{}{}
			}
		}
		dist++
	}
}
