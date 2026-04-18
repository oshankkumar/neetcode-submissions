
type stack []int

func (s *stack) push(val int) {
	*s = append(*s, val)
}

func (s *stack) pop() int {
	last := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return last
}

func (s *stack) top() int {
	return (*s)[len(*s)-1]
}

func (s *stack) len() int {
	return len(*s)
}

var DailyTemperatures = dailyTemperatures

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))

	var stk stack
	for i, temp := range temperatures {
		for stk.len() > 0 && temp > temperatures[stk.top()] {
			idx := stk.pop()
			result[idx] = i - idx
		}
		stk.push(i)
	}

	return result
}
