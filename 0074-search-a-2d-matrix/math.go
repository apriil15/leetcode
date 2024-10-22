package main

// time: O(log(m*n))
// space: O(1)
//
// use [mid/cols][mid%cols] to transfer index to matrix
func searchMatrix_math(matrix [][]int, target int) bool {
	rows := len(matrix)
	cols := len(matrix[0])
	left := 0              // image in 1D
	right := rows*cols - 1 // image in 1D

	for left <= right {
		mid := left + (right-left)/2
		cur := matrix[mid/cols][mid%cols]
		if cur < target {
			left = mid + 1
		} else if target < cur {
			right = mid - 1
		} else {
			return true
		}
	}
	return false
}
