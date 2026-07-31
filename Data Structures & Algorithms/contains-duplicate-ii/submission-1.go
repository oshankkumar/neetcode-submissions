func containsNearbyDuplicate(nums []int, k int) bool {
	seen := make(map[int]int, len(nums))
	for i, n := range nums {
		j, ok := seen[n]
		if !ok {
			seen[n] = i
			continue
		}
		if i-j <= k {
			return true
		}
		seen[n] = i
	}
	return false
}
