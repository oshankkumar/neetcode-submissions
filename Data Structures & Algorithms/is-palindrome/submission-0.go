
func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if !isValidAlphaNumeric(s[i]) {
			i++
			continue
		}
		if !isValidAlphaNumeric(s[j]) {
			j--
			continue
		}
		if toLowerCase(s[i]) != toLowerCase(s[j]) {
			return false
		}
		i++
		j--
	}
	return true
}

func toLowerCase(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return 'a' + (b - 'A')
	}
	return b
}

func isValidAlphaNumeric(b byte) bool {
	return (b >= 'a' && b <= 'z') ||
		(b >= 'A' && b <= 'Z') ||
		(b >= '0' && b <= '9')
}
