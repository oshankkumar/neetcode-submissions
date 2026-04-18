func trap(height []int) int {
	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))

	leftMax[0], rightMax[len(height)-1] = height[0], height[len(height)-1]

	for i := 1; i < len(height); i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	for i := len(height) - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	var maxWater int
	for i, h := range height {
		maxUpperBound := min(leftMax[i], rightMax[i])
		if maxUpperBound > h {
			maxWater += maxUpperBound - h
		}
	}

	return maxWater
}
