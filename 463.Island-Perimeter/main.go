package main

import "fmt"

func main() {
	grid := [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}
	result := islandPerimeter(grid)
	fmt.Println(result)
}

// 4 * 1 的個數 - 2 * 接觸個數
func islandPerimeter(grid [][]int) int {
	oneCount := 0
	horizontalContactCount := 0 // 水平的接觸個數
	verticalContactCount := 0   // 垂直的接觸個數

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				oneCount++
			}
			if j >= 1 && grid[i][j] == 1 && grid[i][j-1] == 1 { // 代表水平相鄰
				horizontalContactCount++
			}
			if i >= 1 && grid[i][j] == 1 && grid[i-1][j] == 1 { // 代表垂直相鄰
				verticalContactCount++
			}
		}
	}
	return 4*oneCount - 2*(verticalContactCount+horizontalContactCount)
}
