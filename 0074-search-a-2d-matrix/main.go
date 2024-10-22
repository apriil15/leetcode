package main

func main() {

}

// time: O(log m + log n)
// space: O(1)
func searchMatrix(matrix [][]int, target int) bool {
	rows := len(matrix)
	cols := len(matrix[0])

	// first binary search to find which row
	top := 0
	bot := rows - 1
	for top <= bot {
		mid := top + (bot-top)/2
		if target < matrix[mid][0] {
			bot = mid - 1
		} else if matrix[mid][cols-1] < target {
			top = mid + 1
		} else {
			break
		}
	}

	if top > bot {
		return false
	}

	// second binary search to find target
	row := top + (bot-top)/2
	left := 0
	right := cols - 1
	for left <= right {
		mid := left + (right-left)/2
		if matrix[row][mid] < target {
			left = mid + 1
		} else if target < matrix[row][mid] {
			right = mid - 1
		} else {
			return true
		}
	}
	return false
}
