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
		x int
		y int
	}
	queue := []point{}
	visit := make(map[point]struct{})

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if rooms[i][j] == 0 {
				queue = append(queue, point{i, j})
				visit[point{i, j}] = struct{}{}
			}
		}
	}

	addRoom := func(i, j int) {
		if i < 0 || j < 0 || i >= rows || j >= cols {
			return
		}
		if _, ok := visit[point{i, j}]; ok {
			return
		}
		if rooms[i][j] == -1 {
			return
		}

		queue = append(queue, point{i, j})
		visit[point{i, j}] = struct{}{}
	}

	dist := 0
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			cur := queue[0]
			queue = queue[1:]

			rooms[cur.x][cur.y] = dist

			addRoom(cur.x+1, cur.y)
			addRoom(cur.x-1, cur.y)
			addRoom(cur.x, cur.y+1)
			addRoom(cur.x, cur.y-1)
		}
		dist++
	}
}
