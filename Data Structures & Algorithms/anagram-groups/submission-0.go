
type Counter[K comparable] map[K]int

func (c Counter[K]) Equal(c2 Counter[K]) bool {
	if len(c) != len(c2) {
		return false
	}

	for k, v := range c {
		if c2[k] != v {
			return false
		}
	}

	return true
}

func NewCounter[K comparable](s []K) Counter[K] {
	c := make(Counter[K])
	for _, v := range s {
		c[v] += 1
	}
	return c
}

func isAnargam(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	c1, c2 := NewCounter([]byte(s1)), NewCounter([]byte(s2))
	return c1.Equal(c2)
}

const removedMarker = "__REMOVED_MARKER__"

func groupAnagrams(strs []string) [][]string {
	var result [][]string
	for i := 0; i < len(strs); i++ {
		if strs[i] == removedMarker {
			continue
		}
		l := []string{strs[i]}
		for j := i + 1; j < len(strs); j++ {
			if !isAnargam(strs[i], strs[j]) {
				continue
			}
			l = append(l, strs[j])
			strs[j] = removedMarker
		}
		result = append(result, l)
	}
	return result
}
