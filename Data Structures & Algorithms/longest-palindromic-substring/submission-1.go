func longestPalindrome(s string) string {
	var start, end int

	palindromeStartsAt := func(left, right int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}
		left++
		right--
		if right-left > end-start {
			start = left
			end = right
		}
	}

	for i := range s {
		palindromeStartsAt(i, i)
		palindromeStartsAt(i, i+1)
	}
	return s[start : end+1]
}
