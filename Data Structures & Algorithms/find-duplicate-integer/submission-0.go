func findDuplicate(nums []int) int {
	counter := make(map[int]int)
	for _, num := range nums {
		counter[num]++
		if counter[num] > 1 {
			return num
		}
	}
	return -1
}

