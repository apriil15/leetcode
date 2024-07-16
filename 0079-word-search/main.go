package main

func main() {

}

// time: O(m * n * 4^len(word))
func exist(board [][]byte, word string) bool {
	rows := len(board)
	columns := len(board[0])

	path := make(map[point]struct{})

	var backtrack func(r, c, i int) bool
	backtrack = func(r, c, i int) bool {
		if i == len(word) { // already traverse entire word
			return true
		}
		if r < 0 || c < 0 ||
			r == rows || c == columns {
			return false
		}
		if board[r][c] != word[i] {
			return false
		}
		if _, ok := path[point{row: r, column: c}]; ok {
			return false
		}

		path[point{row: r, column: c}] = struct{}{}

		// backtrack up, down, left, right
		res := backtrack(r+1, c, i+1) ||
			backtrack(r-1, c, i+1) ||
			backtrack(r, c+1, i+1) ||
			backtrack(r, c-1, i+1)

		delete(path, point{row: r, column: c})

		return res
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if backtrack(i, j, 0) {
				return true
			}
		}
	}
	return false
}

type point struct {
	row    int
	column int
}
