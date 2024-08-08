package main

func main() {

}

func solve(board [][]byte) {
	rows := len(board)
	cols := len(board[0])

	// record valid 'O' in set
	type point struct {
		row int
		col int
	}
	set := make(map[point]struct{})

	var dfs func(i, j int, set map[point]struct{})
	dfs = func(i, j int, set map[point]struct{}) {
		if i < 0 || j < 0 || i >= rows || j >= cols ||
			board[i][j] == 'X' {
			return
		}
		if _, ok := set[point{i, j}]; ok {
			return
		}

		set[point{i, j}] = struct{}{}

		dfs(i+1, j, set)
		dfs(i-1, j, set)
		dfs(i, j+1, set)
		dfs(i, j-1, set)
	}

	// find 'O' from corner
	for i := 0; i < rows; i++ {
		dfs(i, 0, set)
		dfs(i, cols-1, set)
	}
	for j := 0; j < cols; j++ {
		dfs(0, j, set)
		dfs(rows-1, j, set)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if _, ok := set[point{i, j}]; !ok {
				board[i][j] = 'X'
			}
		}
	}
}
