func lengthOfLongestSubstring(s string) int {
    left, right := 0, 0
	charSet := make(map[byte]struct{})
	var maxLen int

	for right < len(s) {
		_, ok := charSet[s[right]]
		if !ok {
			charSet[s[right]] = struct{}{}
			right++
			maxLen = max(maxLen, len(charSet))
		} else {
			delete(charSet, s[left])
			left++
		}
	}

	return maxLen
}
