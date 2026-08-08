
func longestCommonSubsequence(text1 string, text2 string) int {
	memo := make([][]int, len(text1))
	for i := range memo {
		memo[i] = make([]int, len(text2))
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	return findLongestCommonSubsequence(0, 0, memo, text1, text2)
}

func findLongestCommonSubsequence(i, j int, memo [][]int, text1 string, text2 string) int {
	if i >= len(text1) || j >= len(text2) {
		return 0
	}

	if memo[i][j] != -1 {
		return memo[i][j]
	}

	if text1[i] == text2[j] {
		return 1 + findLongestCommonSubsequence(i+1, j+1, memo, text1, text2)
	}

	memo[i][j] = max(
		findLongestCommonSubsequence(i+1, j, memo, text1, text2),
		findLongestCommonSubsequence(i, j+1, memo, text1, text2),
	)
	return memo[i][j]
}
