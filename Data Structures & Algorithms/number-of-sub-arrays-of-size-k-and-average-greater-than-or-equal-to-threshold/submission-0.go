func numOfSubarrays(arr []int, k int, threshold int) int {
	var l, r int
	var currSum int

	var count int
	for r < len(arr) {
		if r-l+1 < k {
			currSum += arr[r]
			r++
			continue
		}

		currSum += arr[r]
		if currSum/k >= threshold {
			count++
		}

		currSum -= arr[l]
		l++
		r++
	}
	return count
}
