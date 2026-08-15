func spiralOrder(matrix [][]int) []int {
	var result []int
	top, bottom := 0, len(matrix)
	left, right := 0, len(matrix[0])

	for top < bottom && left < right {
		for j := left; j < right; j++ {
			result = append(result, matrix[top][j])
		}
		top++

		for i := top; i < bottom; i++ {
			result = append(result, matrix[i][right-1])
		}
		right--

		if !(left < right && top < bottom) {
			break
		}

		for j := right-1; j >= left; j-- {
			result = append(result, matrix[bottom-1][j])
		}
		bottom--

		for i := bottom-1; i >= top; i-- {
			result = append(result, matrix[i][left])
		}
		left++
	}

	return result
}
