
func characterReplacement(s string, k int) int {
	var left, right, maxLen, maxFreq int

	counter := make(map[byte]int)

	for right < len(s) {
		counter[s[right]]++
		maxFreq = max(maxFreq, counter[s[right]])

		if (right-left+1)-maxFreq > k {
			counter[s[left]]--
            left++
		}
		
		maxLen = max(maxLen, right-left+1)
		right++
	}
	return maxLen
}
