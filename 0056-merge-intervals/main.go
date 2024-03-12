package main

import (
	"fmt"
	"sort"
)

func main() {
	a := [][]int{{1, 3}, {0, 2}, {4, 5}, {2, 2}}
	fmt.Println(merge(a))
}

// time: O(n logn)
// space: O(n)
func merge(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{}

	i := 0
	for i <= len(intervals)-2 {
		a := intervals[i]
		b := intervals[i+1]
		if isOverlap(a, b) {
			// replace the latter one
			intervals[i+1] = mergeInterval(a, b)
		} else {
			result = append(result, a)
		}

		i++
	}

	return append(result, intervals[len(intervals)-1])
}

func isOverlap(a, b []int) bool {
	return a[0] <= b[1] && b[0] <= a[1]
}

func mergeInterval(a, b []int) []int {
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
