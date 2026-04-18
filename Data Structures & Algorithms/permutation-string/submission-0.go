
import "maps"

func checkInclusion(s1 string, s2 string) bool {
	windowSize := len(s1)
	var left, right int

	counter := make(map[byte]int)
	for i := range s1 {
		counter[s1[i]]++
	}

	seen := make(map[byte]int)
	for right = 0; right < len(s2); right++ {
		seen[s2[right]]++
		if right-left+1 < windowSize {
			continue
		}
		if maps.Equal(counter, seen) {
			return true
		}

		seen[s2[left]]--
        if seen[s2[left]] == 0 {
			delete(seen, s2[left])
		}
		left++
	}
	return false
}