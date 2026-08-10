

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var b strings.Builder
	for _, str := range strs {
		b.WriteRune('#')
		fmt.Fprintf(&b, "%03d", len(str))
		b.WriteString(str)
	}
	return b.String()
}

func (s *Solution) Decode(encoded string) []string {
	var strs []string
	var i int
	for i < len(encoded) {
		var strLen int
		if encoded[i] == '#' {
			for range 3 {
				i++
				strLen = strLen*10 + int(encoded[i]-'0')
			}
		}

		i++
		strs = append(strs, encoded[i:strLen+i])
		i += strLen
	}
	return strs
}
