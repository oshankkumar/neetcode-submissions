
func lengthOfLongestSubstring(s string) int {
	var l, r int

	seen := make(map[byte]struct{})
	var maxLen int
	for r < len(s) {
		if _, ok := seen[s[r]]; !ok {
			seen[s[r]] = struct{}{}
			maxLen = max(maxLen, len(seen))
			r++
			continue
		}
		delete(seen, s[l])
		l++
	}

	return maxLen
}
