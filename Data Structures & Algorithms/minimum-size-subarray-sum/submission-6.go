func minSubArrayLen(target int, nums []int) int {
	minLen := math.MaxInt
	var l int

	var currSum int
	for r := 0; r < len(nums); r++ {
		currSum += nums[r]

		if currSum < target {
			continue
		}

		for currSum >= target {
			minLen = min(minLen, r-l+1)
			currSum -= nums[l]
			l++
		}
	}

	if minLen == math.MaxInt {
		return 0
	}

	return minLen
}
