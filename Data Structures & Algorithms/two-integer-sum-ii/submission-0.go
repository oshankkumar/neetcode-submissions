func twoSum(numbers []int, target int) []int {
    i, j := 0, len(numbers)-1
	for i < j {
		currSum := numbers[i] + numbers[j]
		switch {
		case currSum == target:
			return []int{i + 1, j + 1}
		case currSum < target:
			i++
		case currSum > target:
			j--
		}
	}

	return nil
}
