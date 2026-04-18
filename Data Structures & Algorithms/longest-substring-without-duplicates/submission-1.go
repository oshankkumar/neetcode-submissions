func lengthOfLongestSubstring(s string) int {
    left, right := 0, -1
	charSet := make(map[byte]struct{})
	var maxLen int

	for right < len(s)-1 {
		_, ok := charSet[s[right+1]]
		if !ok {
			charSet[s[right+1]] = struct{}{}
			right++
			maxLen = max(maxLen, len(charSet))
		} else {
			delete(charSet, s[left])
			left++
		}
	}

	return maxLen
}
