func twoSum(nums []int, target int) []int {
	positions := make(map[int]int)

	for i, n := range nums {
		v, ok := positions[n]
		if !ok {
			positions[n] = i
			continue
		}
		if v > len(nums) {
			continue
		}

		positions[n] += (i + len(nums))

	}

	for i, n := range nums {
		pos, ok := positions[target-n]
		if !ok || pos == i {
			continue
		}
		if pos < len(nums) {
			return []int{i, pos}
		}
		return []int{i, pos - len(nums) - i}
	}

	return []int{-1, -1}
}
