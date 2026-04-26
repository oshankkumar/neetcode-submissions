func lengthOfLongestSubstring(s string) int {
	var left, right int

	str := []rune(s)
	charSeen := make(map[rune]struct{})
	var maxLen int
	for right < len(s) {
		curr := str[right]
		if _, ok := charSeen[curr]; ok {
			delete(charSeen, str[left])
			left++
			continue
		}
		charSeen[curr] = struct{}{}
		maxLen = max(maxLen, len(charSeen))
		right++

	}
	return maxLen
}
