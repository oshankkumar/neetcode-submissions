
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	memo := make([][]int, len(obstacleGrid))
	for i := range memo {
		memo[i] = make([]int, len(obstacleGrid[0]))
	}
	return findUniquePathsWithObstacles(0, 0, obstacleGrid, memo)
}

func findUniquePathsWithObstacles(i, j int, obstacleGrid [][]int, memo [][]int) int {
	if i >= len(obstacleGrid) || j >= len(obstacleGrid[0]) {
		return 0
	}

	if obstacleGrid[i][j] == 1 {
		return 0
	}

	if i == len(obstacleGrid)-1 && j == len(obstacleGrid[0])-1 {
		return 1
	}

	if memo[i][j] != 0 {
		return memo[i][j]
	}

	count := findUniquePathsWithObstacles(i+1, j, obstacleGrid, memo) + findUniquePathsWithObstacles(i, j+1, obstacleGrid, memo)
	
	memo[i][j] = count

	return count
}
