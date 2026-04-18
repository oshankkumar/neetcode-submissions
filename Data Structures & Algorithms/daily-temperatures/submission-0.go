func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		var after int
		for j := i; j < len(temperatures); j++ {
			if temperatures[j] > temperatures[i] {
				after = j - i
				break
			}
		}
		result[i] = after
	}
	return result
}