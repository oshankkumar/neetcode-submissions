
func maxArea(heights []int) int {
	l, r := 0, len(heights)-1
	var mArea int
	for l < r {
		mArea = max(mArea, (r-l)*min(heights[l], heights[r]))
		if heights[l] > heights[r] {
			r--
		} else {
			l++
		}
	}
	return mArea
}
