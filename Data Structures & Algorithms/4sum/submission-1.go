
func fourSum(nums []int, target int) [][]int {
	var result [][]int

	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left, right := j+1, len(nums)-1

			required := target - (nums[i] + nums[j])

			for left < right {
				if nums[left]+nums[right] < required {
					left++
					continue
				}

				if nums[left]+nums[right] > required {
					right--
					continue
				}

				result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}

				for left < right && nums[right] == nums[right+1] {
					right--
				}

			}
		}
	}
	return result
}
