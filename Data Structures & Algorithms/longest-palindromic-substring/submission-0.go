func longestPalindrome(s string) string {
	var start, end int

	for i := range s {
		left, right := LongestPalindromIndexStartsAt(s, i, i)
		if right-left > end-start {
			start = left
			end = right
		}
		left, right = LongestPalindromIndexStartsAt(s, i, i+1)
		if right-left > end-start {
			start = left
			end = right
		}
	}
	return s[start : end+1]
}

func LongestPalindromIndexStartsAt(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right - 1
}
