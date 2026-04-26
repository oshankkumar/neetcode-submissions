
func rob(nums []int) int {
	memo := make([]int, len(nums))
	for i := range memo {
		memo[i] = -1
	}
	return robHouse(nums, 0, memo)
}

func robHouse(nums []int, left int, memo []int) int {
	if left == len(nums)-1 {
		return nums[left]
	}

	if left == len(nums)-2 {
		return max(nums[left], nums[left+1])
	}

	if memo[left] != -1 {
		return memo[left]
	}

	if memo[left+1] == 1 {
		memo[left+1] = robHouse(nums, left+1, memo)
	}

	memo[left] = max(
		nums[left]+robHouse(nums, left+2, memo),
		robHouse(nums, left+1, memo),
	)

	return memo[left]
}
