
func countBits(n int) []int {
	result := make([]int, 0)
	for i := range n + 1 {
		result = append(result, countBitsSingle(i))
	}
	return result
}

func countBitsSingle(n int) int {
	var count int
	for range 10 {
		if n&1 == 1 {
			count++
		}
		n >>= 1
	}
	return count
}
