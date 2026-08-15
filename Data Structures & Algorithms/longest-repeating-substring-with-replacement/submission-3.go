
func characterReplacement(s string, k int) int {
	var l, r int
	freq := make(map[byte]int)
	var (
		maxFreqSeen int
		maxLen      int
	)

	for r < len(s) {
		freq[s[r]]++
		maxFreqSeen = max(maxFreqSeen, freq[s[r]])

		windowLen := (r - l + 1)

		if windowLen-maxFreqSeen <= k {
			maxLen = max(maxLen, windowLen)
			r++
			continue
		}

		freq[s[l]]--
		l++
		r++
	}

	return maxLen
}
