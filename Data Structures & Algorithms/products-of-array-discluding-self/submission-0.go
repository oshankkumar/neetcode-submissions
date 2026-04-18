func productExceptSelf(nums []int) []int {
	nonZeroProduct := 1
	zeroAt := -1

	out := make([]int, len(nums))

	for i, num := range nums {
		if num != 0 {
			nonZeroProduct *= num
			continue
		}
		if zeroAt > -1 {
			return out
		}
		zeroAt = i
	}

	for i, num := range nums {
		if num == 0 {
			out[i] = nonZeroProduct
		} else if zeroAt > -1 {
			out[i] = 0
		} else {
			out[i] = nonZeroProduct / num
		}
	}

	return out
}