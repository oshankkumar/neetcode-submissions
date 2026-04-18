func hammingWeight(n int) int {
	var count int
	for range 32 {
		if n&1 == 1 {
			count++
		}
		n >>= 1
	}
	return count
}
