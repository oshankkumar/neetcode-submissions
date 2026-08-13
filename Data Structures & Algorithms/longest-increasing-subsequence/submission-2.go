func LengthOfLIS(nums []int, i int, memo []int) int {
	if i == len(nums)-1 {
		return 1
	}

	if memo[i] != 0 {
		return memo[i]
	}

	maxLen := 1
	for j := i + 1; j < len(nums); j++ {
		if nums[j] > nums[i] {
			maxLen = max(maxLen, 1+LengthOfLIS(nums, j, memo))
		}
	}

	memo[i] = maxLen

	return maxLen
}

func lengthOfLIS(nums []int) int {
	memo := make([]int, len(nums))
	maxLen := 1
	for i := range nums {
		maxLen = max(maxLen, LengthOfLIS(nums, i, memo))
	}
	return maxLen
}
