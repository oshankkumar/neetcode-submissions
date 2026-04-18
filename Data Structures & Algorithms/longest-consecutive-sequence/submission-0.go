func longestConsecutive(nums []int) int {
	searchIndex := make(map[int]int)
	for i, num := range nums {
		searchIndex[num] = i
	}

	var longestLen int
	for _, num := range nums {
		_, ok := searchIndex[num-1]
		if ok {
			continue
		}

		nextIdx, nextIdxExists := searchIndex[num+1]
		currLen := 1
		for nextIdxExists {
			currLen++
			nextIdx, nextIdxExists = searchIndex[nums[nextIdx]+1]
		}

		longestLen = max(longestLen, currLen)
	}

	return longestLen
}
