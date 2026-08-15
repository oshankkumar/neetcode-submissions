import "maps"

func checkInclusion(s1 string, s2 string) bool {
	s1BytesFreq := make(map[byte]int)

	for i := range s1 {
		s1BytesFreq[s1[i]]++
	}

	var l int
	currentFreq := make(map[byte]int)
	for r := 0; r < len(s2); r++ {
		currentFreq[s2[r]]++
		if r-l+1 < len(s1) {
			continue
		}

		if maps.Equal(currentFreq, s1BytesFreq) {
			return true
		}

		currentFreq[s2[l]]--
		if currentFreq[s2[l]] == 0 {
			delete(currentFreq, s2[l])
		}
		l++
	}

	return false
}