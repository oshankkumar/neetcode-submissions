func canJump(nums []int) bool {
	memo := make(map[int]bool,len(nums))
    return canReachEndFrom(0,memo,nums)
}

func canReachEndFrom(i int,memo map[int]bool,nums []int) bool {
    if i == len(nums) -1 {
		return true
	}

	if i >= len(nums) || nums[i] == 0 {
		return false
	}

	if val,ok := memo[i];ok {
		return val
	}

	for j := 1; j <= nums[i];j++  {
		if canReachEndFrom(i+j,memo,nums) {
			memo[i] = true 
			return true
		}
	}

	memo[i] = false
	return false
}