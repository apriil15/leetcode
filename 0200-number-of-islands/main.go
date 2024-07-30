package main

func main() {

}

func numIslands(grid [][]byte) int {
	rows := len(grid)
	cols := len(grid[0])

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || j < 0 || i >= rows || j >= cols {
			return
		}
		if grid[i][j] == '0' {
			return
		}

		// if it is '1' (land), then set '0' as walked
		grid[i][j] = '0'

		// up, down, left, right
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}

	var res int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				res++
				dfs(i, j)
			}
		}
	}

	return res
}
