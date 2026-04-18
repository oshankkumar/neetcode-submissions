func maxArea(heights []int) int {
	l, r := 0, len(heights)-1
	var maxArea int
	for l < r {
		maxArea = max(
			maxArea,
			(r-l)*min(heights[l], heights[r]),
		)
		if heights[l] < heights[r] {
			l++
		} else {
			r--
		}
	}
	return maxArea
}
