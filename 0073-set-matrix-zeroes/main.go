package main

func main() {

}

// time: O(m*n)
// space: O(1)
func setZeroes(matrix [][]int) {
	rows := len(matrix)
	columns := len(matrix[0])

	// whether row zero should be set 0
	rowZero := false

	// if there is any value which is 0 at row zero, set rowZero to be true
	// mark first row's and column's value as 0 when the value is 0
	// if first row's and column's value is 0, then the whole row or column should be set 0
	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			if matrix[r][c] == 0 {
				if r == 0 {
					rowZero = true
				} else {
					matrix[r][0] = 0
					matrix[0][c] = 0
				}
			}
		}
	}

	for r := 1; r < rows; r++ {
		for c := 1; c < columns; c++ {
			if matrix[0][c] == 0 || matrix[r][0] == 0 {
				matrix[r][c] = 0
			}
		}
	}

	// update column zero
	if matrix[0][0] == 0 {
		for r := 0; r < rows; r++ {
			matrix[r][0] = 0
		}
	}
	// update row zero
	if rowZero {
		for c := 0; c < columns; c++ {
			matrix[0][c] = 0
		}
	}
}

// time: O(m*n)
// space: O(m+n)
func setZeroes_2(matrix [][]int) {
	rows := len(matrix)
	columns := len(matrix[0])

	// find 0 position
	rowHasZero := make(map[int]struct{})
	columnHasZero := make(map[int]struct{})
	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			if matrix[r][c] == 0 {
				rowHasZero[r] = struct{}{}
				columnHasZero[c] = struct{}{}
			}
		}
	}

	for r := range rowHasZero {
		for c := 0; c < columns; c++ {
			matrix[r][c] = 0
		}
	}

	for c := range columnHasZero {
		for r := 0; r < rows; r++ {
			matrix[r][c] = 0
		}
	}
}

// time: O(m*n)
// space: O(0's count)
func setZeroes_3(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])

	// find 0 position
	zeroPositions := [][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				zeroPositions = append(zeroPositions, []int{i, j})
			}
		}
	}

	// modify to 0
	for _, position := range zeroPositions {
		x := position[0]
		y := position[1]

		// modify horizontal
		for i := 0; i < n; i++ {
			matrix[x][i] = 0
		}

		// modify vertical
		for i := 0; i < m; i++ {
			matrix[i][y] = 0
		}
	}
}
