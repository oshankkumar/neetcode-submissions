func lengthOfLIS(nums []int) int {
	lis := make([]int, len(nums))
	for i := range lis {
		lis[i] = 1
	}

	for i := range nums {
		for j := i - 1; j >= 0; j-- {
			if nums[j] < nums[i] {
				lis[i] = max(lis[i], 1+lis[j])
			}
		}
	}

	maxLen := lis[0]
	for _, l := range lis[1:] {
		maxLen = max(maxLen, l)
	}
	return maxLen
}
