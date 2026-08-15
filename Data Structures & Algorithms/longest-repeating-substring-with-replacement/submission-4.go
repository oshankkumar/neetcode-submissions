
func characterReplacement(s string, k int) int {
	var l int
	freq := make(map[byte]int)
	var (
		maxFreqSeen int
		maxLen      int
	)

	for r := 0; r < len(s);r++ {
		freq[s[r]]++
		maxFreqSeen = max(maxFreqSeen, freq[s[r]])

		windowLen := (r - l + 1)

		if windowLen-maxFreqSeen <= k {
			maxLen = max(maxLen, windowLen)
			continue
		}

		freq[s[l]]--
		l++
	}

	return maxLen
}
