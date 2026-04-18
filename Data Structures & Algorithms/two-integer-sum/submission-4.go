
func twoSum(nums []int, target int) []int {
	positions := make(map[int]int)

	for i, n := range nums {
		j, ok := positions[target-n]
		if ok && i != j {
			return []int{j, i}
		}
		positions[n] = i
	}

	return []int{-1, -1}
}
