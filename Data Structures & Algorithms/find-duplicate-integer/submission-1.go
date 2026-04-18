func findDuplicate(nums []int) int {
	var i int
	for i < len(nums) {
		if nums[i] == i+1 {
			i++
			continue
		}

		if nums[nums[i]-1] == nums[i] {
			return nums[i]
		}

		nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
	}

	return -1
}

