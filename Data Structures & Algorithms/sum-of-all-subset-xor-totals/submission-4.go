
func subsetXORSum(nums []int) int {
	subsets := FindAllSubsets(nums)
	var sum int
	for _, subset := range subsets {
		var xor int
		for _, num := range subset {
			xor ^= num
		}
		sum += xor
	}
	return sum
}

func FindAllSubsets(num []int) [][]int {
	if len(num) == 0 {
		return [][]int{num}
	}

	subsets := FindAllSubsets(num[1:])

	var additional [][]int
	for _, subset := range subsets {
		additional = append(additional, append([]int{num[0]}, subset...))
	}

	return append(subsets, additional...)
}
