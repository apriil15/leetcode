package main

import "strconv"

func main() {

}

// time: O(9*9)
// space: O(9*3)
func isValidSudoku(board [][]byte) bool {
	rowToDigits := make(map[int]map[byte]struct{})
	colToDigits := make(map[int]map[byte]struct{})
	// key: "00", "01", ..., "22"
	squareToDigits := make(map[string]map[byte]struct{})

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			digit := board[i][j]
			if digit == '.' {
				continue
			}

			squareKey := strconv.Itoa(i/3) + strconv.Itoa(j/3)

			// initialize
			if _, ok := rowToDigits[i]; !ok {
				rowToDigits[i] = make(map[byte]struct{})
			}
			if _, ok := colToDigits[j]; !ok {
				colToDigits[j] = make(map[byte]struct{})
			}
			if _, ok := squareToDigits[squareKey]; !ok {
				squareToDigits[squareKey] = make(map[byte]struct{})
			}

			// check exist
			if _, ok := rowToDigits[i][digit]; ok {
				return false
			}
			if _, ok := colToDigits[j][digit]; ok {
				return false
			}
			if _, ok := squareToDigits[squareKey][digit]; ok {
				return false
			}

			rowToDigits[i][digit] = struct{}{}
			colToDigits[j][digit] = struct{}{}
			squareToDigits[squareKey][digit] = struct{}{}
		}
	}
	return true
}
