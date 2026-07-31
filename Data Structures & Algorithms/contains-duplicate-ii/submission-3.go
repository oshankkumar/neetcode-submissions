func containsNearbyDuplicate(nums []int, k int) bool {
	window := make(map[int]struct{})
	var l int
	for r, num := range nums {
		if r-l > k {
			delete(window, nums[l])
			l++
		}
		if _, ok := window[num]; ok {
			return true
		}
		window[num] = struct{}{}
	}
	return false
}