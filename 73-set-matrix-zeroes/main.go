package main

func main() {

}

func setZeroes(matrix [][]int) {
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
