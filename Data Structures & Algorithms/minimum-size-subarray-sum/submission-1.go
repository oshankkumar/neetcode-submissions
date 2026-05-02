
func minSubArrayLen(target int, nums []int) int {
	left, right := 0, 0

	minLen := len(nums) + 1

	runningSum := nums[right]
	for left <= right {
		if runningSum < target {
			right++
			if right >= len(nums) {
				break
			}
			runningSum += nums[right]
			continue
		}

		minLen = min(minLen, right-left+1)
		runningSum -= nums[left]
		left++
	}

	if minLen == len(nums)+1 {
		return 0
	}

	return minLen
}
