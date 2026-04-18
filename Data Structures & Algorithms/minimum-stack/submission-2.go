type MinStack struct {
	vals    []int
	currMin int
}

func Constructor() MinStack {
	return MinStack{currMin: math.MaxInt}
}

func (m *MinStack) Push(val int) {
	if len(m.vals) == 0 {
		m.currMin = val
		m.vals = append(m.vals, 0)
		return
	}

	m.vals = append(m.vals, val-m.currMin)
	m.currMin = min(m.currMin, val)
}

func (m *MinStack) Pop() {
	if len(m.vals) == 0 {
		return
	}

	top := m.vals[len(m.vals)-1]
	if top < 0 {
		m.currMin = m.currMin - top
	}
	m.vals = m.vals[:len(m.vals)-1]
}

func (m *MinStack) Top() int {
	top := m.vals[len(m.vals)-1]
	if top > 0 {
		return top + m.currMin
	}
	return m.currMin
}

func (m *MinStack) GetMin() int {
	return m.currMin
}
