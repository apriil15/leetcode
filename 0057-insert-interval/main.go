package main

import "fmt"

func main() {
	fmt.Println(insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))
}

// time: O(n)
// space: O(n)
func insert(intervals [][]int, newInterval []int) [][]int {
	result := make([][]int, 0, len(intervals))
	for i := 0; i < len(intervals); i++ {
		itv := intervals[i]

		// newInterval is smaller than all the other intervals right side
		if newInterval[1] < itv[0] {
			result = append(result, newInterval)
			return append(result, intervals[i:]...)
		}

		if itv[1] < newInterval[0] {
			result = append(result, itv)
		} else {
			newInterval = merge(newInterval, itv) // overlap
		}
	}

	return append(result, newInterval)
}

func isOverlap(a, b []int) bool {
	return a[0] <= b[1] && b[0] <= a[1]
}

func merge(a, b []int) []int {
	return []int{min(a[0], b[0]), max(a[1], b[1])}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
