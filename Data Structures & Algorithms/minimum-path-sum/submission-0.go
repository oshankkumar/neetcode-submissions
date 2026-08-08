
func minPathSum(grid [][]int) int {
	memo := make([][]int, len(grid))
	for i := range memo {
		memo[i] = make([]int, len(grid[0]))
	}
	return findMinPathSum(0, 0, memo, grid)
}

func findMinPathSum(i, j int, memo [][]int, grid [][]int) int {
	if i == len(grid)-1 && j == len(grid[0])-1 {
		return grid[i][j]
	}

	if i >= len(grid) || j >= len(grid[0]) {
		return math.MaxInt
	}

	if val := memo[i][j]; val > 0 {
		return val
	}

	memo[i][j] = grid[i][j] + min(findMinPathSum(i+1, j, memo, grid), findMinPathSum(i, j+1, memo, grid))
	return memo[i][j]
}
