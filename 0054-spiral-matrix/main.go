package main

import "fmt"

func main() {
	result := spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	fmt.Println(result)
}

// time: O(m*n)
// space: O(1)
func spiralOrder(matrix [][]int) []int {
	rows := len(matrix)
	columns := len(matrix[0])
	all := rows * columns
	result := make([]int, 0, rows*columns)

	rowBegin := 0
	rowEnd := rows - 1
	columnBegin := 0
	columnEnd := columns - 1
	for {
		// right
		for c := columnBegin; c <= columnEnd; c++ {
			result = append(result, matrix[rowBegin][c])
			if len(result) == all {
				return result
			}
		}
		rowBegin++

		// down
		for r := rowBegin; r <= rowEnd; r++ {
			result = append(result, matrix[r][columnEnd])
			if len(result) == all {
				return result
			}
		}
		columnEnd--

		// left
		for c := columnEnd; c >= columnBegin; c-- {
			result = append(result, matrix[rowEnd][c])
			if len(result) == all {
				return result
			}
		}
		rowEnd--

		// up
		for r := rowEnd; r >= rowBegin; r-- {
			result = append(result, matrix[r][columnBegin])
			if len(result) == all {
				return result
			}
		}
		columnBegin++
	}
}
